package queue

import (
	"encoding/json"
	"github.com/shellus/pkg/logs"
)

// go的任务队列，或者说goroutine管理器

//const redis_prefix string = "queue"

type queue struct {
	channelName string
	jobs       chan *Job
	subscriber  func(j *Job) (err error)
	thrDispathChan chan interface{}
}

type Job struct {
	retry int // 已经尝试次数
	Value interface{} // 任务参数
	ThrValue interface{} // 线程上下文
}

func NewQueue() (q *queue) {
	q = &queue{
		jobs: make(chan *Job, 10000),
		thrDispathChan: make(chan interface{}),
	}
	go func() {
		for thrValue := range q.thrDispathChan{
			go q.thrDispath(thrValue)
		}
	}()

	return
}

func (q *queue) AddThread(value interface{}){
	q.thrDispathChan <- value
}
func (q *queue) AddEmptyThread(count int){
	for i := 0; i < count; i++{
		q.AddThread(true)
	}
}
func (q *queue) Close() {
	close(q.thrDispathChan)
}

func (q *queue) thrDispath (thrValue interface{}) {
	for j := range q.jobs {
		j.ThrValue = thrValue
		q.call(j)
	}
}

func (q *queue) Pub(j *Job) {
	j.retry++
	q.jobs <- j
}

func (q *queue) Sub(f func(j *Job) (err error)) {
	q.subscriber = f
}

func (q *queue) call(j *Job) {
	// call
	err := q.subscriber(j)

	if err != nil {
		logs.Warning("queue call error: %s", err)
		if j.retry < 3 {
			q.Pub(j)
		}else {
			// todo 如何优雅的丢弃job
			logs.Error("queue call error has max retry: %s", err)
		}
	}
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