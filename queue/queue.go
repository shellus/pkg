package queue

import (
	"context"
	"encoding/json"
	"fmt"
)

// go的任务队列，或者说goroutine管理器

//const redis_prefix string = "queue"

type Call func(j *Job) (err error)

type thr struct {
	index uint
	value interface{} // （线程用户值）
	ctx   context.Context
}

type queue struct {
	channelName string
	jobs        chan *Job
	thrChan     chan *thr
	thrExitChan chan *thr // 如果线程要退出，需要往这个chan里推数据，
	thrNum      uint
	f           Call // 回调函数 （消费函数）
}

type Job struct {
	retry int         // 已经尝试次数
	Value interface{} // 任务参数
}

func NewQueue(f Call) *queue {
	q := &queue{
		jobs:        make(chan *Job, 10000),
		thrChan:     make(chan *thr),
		thrExitChan: make(chan *thr),
		f:           f,
	}

	// 等待新增工作线程信号，直到调用queue.Close()
	go q.thrDispatch()

	return q
}

func (q *queue) AddThread(count int) {
	for i := 0; i < count; i++ {
		q.thrChan <- &thr{index: q.thrNum + 1}
	}
}

// 所有工作线程执行完手头工作就退出
func (q *queue) Close() {
	close(q.thrChan)
	close(q.jobs)
}
func (q *queue) thrDispatch() {
	for thr := range q.thrChan {
		go q.jobDispatch(thr)
		q.thrNum = q.thrNum + 1
	}
}

func (q *queue) jobDispatch(thr *thr) {

A:
	for {
		select {
		// 如果接受我退出那我就退出
		case q.thrExitChan <- thr:
			break A
		// 如果有job我就执行job
		case j := <-q.jobs:
			q.call(j)
		}
	}
	q.thrNum = q.thrNum - 1
}
func (q *queue) call(j *Job) {
	err := q.f(j)

	if err != nil {
		fmt.Printf("queue call error: %s", err)
		if j.retry < 3 {
			q.Pub(j)
		} else {
			// todo 如何优雅的丢弃job
			fmt.Printf("queue call error has max retry: %s", err)
		}
	}
}

func (q *queue) Pub(j *Job) {
	j.retry++
	q.jobs <- j
}

func serialization(j *Job) (s string) {
	buf, err := json.Marshal(j)
	if err != nil {
		return ""
	}
	return string(buf)
}

func deserialization(s string, j *Job) (err error) {
	err = json.Unmarshal([]byte(s), j)
	if err != nil {
		return err
	}
	return nil
}
