package queue

import (
	"fmt"
	"encoding/json"
)



// go的任务队列，或者说goroutine管理器

//const redis_prefix string = "queue"

type queue struct {
	channelName string
	jobs       chan *Job
	concurrent  chan bool
	subscriber  func(j *Job) (err error)
	exit bool
}

type Job struct {
	retry int
	Value interface{}
}

func NewQueue(concurrentNumber int, channelName string) (q *queue) {
	q = &queue{
		channelName: channelName,
		concurrent: make(chan bool, concurrentNumber),
		jobs: make(chan *Job, 10000),
		exit: false,
	}
	return
}


func (q *queue) Pub(j *Job) {
	j.retry++
	q.jobs <- j
}

func (q *queue) Sub(f func(j *Job) (err error)) {
	q.subscriber = f
}

/**
运行完所有job即退出
 */
func (q *queue) WorkDoneQuit() {
	L:
	for {
		select {
		case j := <-q.jobs:
			q.concurrent <- true
			go q.call(j)
		default:
			for i := 0; i < cap(q.concurrent); i++ {
				q.concurrent <- true
			}
			break L
		}
	}
}

func (q *queue) call(j *Job) {
	defer func() {
		<-q.concurrent
	}()

	// call
	err := q.subscriber(j)

	if err != nil {

		// todo 如何log
		fmt.Printf("%+v\n", err)

		if j.retry < 3 {
			q.Pub(j)
		}else {
			// todo 如何优雅的丢弃job
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