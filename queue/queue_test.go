package queue

import (
	"fmt"
	"testing"
)
var q = NewQueue(10, "goTest")

func TestQueue(t *testing.T) {

	q.Sub(func(j *Job)(err error) {
		fmt.Println(j.Value)
		return nil
	})

	q.Pub(&Job{Value:"hahaha"})
	q.Pub(&Job{Value:"hahaha2"})
	q.Pub(&Job{Value:"hahaha3"})
	q.Pub(&Job{Value:"hahaha4"})

	q.WorkDoneQuit()
}


