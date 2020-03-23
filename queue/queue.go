package queue

import (
	"context"
	"fmt"
)

// go的任务队列

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
	ChannelName string
	Retry       int         // 已经尝试次数
	Value       interface{} // 任务参数
}

func NewQueue(f Call, channelName string) *queue {
	if channelName == "" {
		channelName = "default"
	}
	q := &queue{
		channelName: channelName,
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
		q.thrNum = q.thrNum + 1
	}
}
func (q *queue) SubThread(count uint) error {
	if count <= 0 {
		return fmt.Errorf("SubThread count(%d) <= 0", count, q.thrNum)
	}
	if count > q.thrNum {
		return fmt.Errorf("SubThread count(%d) > queue.thrNum(%d)", count, q.thrNum)
	}
	q.thrNum = q.thrNum - count

	// 消费掉thrExitChan，以让工作线程退出
	for i := 0; i < int(count); i++ {
		<-q.thrExitChan
	}

	return nil
}

// 所有工作线程执行完手头工作就退出
func (q *queue) Close() {
	close(q.thrChan)
	close(q.jobs)
}
func (q *queue) thrDispatch() {
	for thr := range q.thrChan {
		go q.jobDispatch(thr)
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
}
func (q *queue) call(j *Job) {
	err := q.f(j)

	if err != nil {
		fmt.Printf("queue call error: %s", err)
		if j.Retry < 3 {
			q.pubJob(j)
		} else {
			// todo 如何优雅的丢弃job
			fmt.Printf("queue call error has max Retry: %s", err)
		}
	}
}

func (q *queue) Pub(value interface{}, channelName string) {
	if channelName == "default" {
		channelName = "default"
	}

	q.pubJob(&Job{Value: value, ChannelName: channelName})
}
func (q *queue) pubJob(j *Job) {
	q.jobs <- j
}
