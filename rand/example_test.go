package rand_test

import (
	"testing"
	"github.com/shellus/pkg/rand"
)

func TestInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(rand.Int64(int64(500),int64(1000)))
	}
}

func TestInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(rand.Int(500,1000))
	}
}