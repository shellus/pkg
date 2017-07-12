package queue

import (
	"fmt"
	"testing"
	"time"
)


func TestQueue(t *testing.T) {
	var q = NewQueue()

	q.AddThread("A")
	q.AddThread("B")
	q.AddThread("C")

	handle := func(j *Job)(err error) {

		fmt.Println(j.ThrValue, j.Value)

		time.Sleep(time.Millisecond * 100)

		return nil
	}

	q.Sub(handle)



	q.Pub(&Job{Value:"hahaha"})
	q.Pub(&Job{Value:"hahaha2"})
	q.Pub(&Job{Value:"hahaha3"})
	q.Pub(&Job{Value:"hahaha4"})
	q.Pub(&Job{Value:"hahaha5"})
	q.Pub(&Job{Value:"hahaha6"})
	q.Pub(&Job{Value:"hahaha7"})

	for{
		time.Sleep(time.Second)
	}
}


