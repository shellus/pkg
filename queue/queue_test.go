package queue_test

import (
	"../queue"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

//type Item struct {
//	name string
//}
//
//type ItemQueue *queue
//
//func (j *queue) addItem(item Item) {
//	j.AddContextThread()
//}

func TestQueue(t *testing.T) {

	wg := sync.WaitGroup{}

	handle := func(j *queue.Job) (err error) {
		time.Sleep(time.Second)
		fmt.Println(j.Value)

		time.Sleep(time.Millisecond * 100)

		wg.Done()

		return nil
	}

	var q = queue.NewQueue(handle, "")
	q.AddThread(3)

	testNum := 10

	wg.Add(testNum)

	for i := 1; i <= testNum; i++ {
		q.Pub("hahaha"+strconv.Itoa(i), "")
	}

	wg.Wait()
}
