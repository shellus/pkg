package sutil_test

import (
	"testing"
	"github.com/shellus/pkg/sutil"
	"path/filepath"
)

func TestHomeDir(t *testing.T) {
	t.Log(filepath.Join(sutil.HomeDir(),".ssh","id_rsa.pub"))

}

func TestInt64(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(sutil.RandInt64(int64(500),int64(1000)))
	}
}

func TestInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(sutil.RandInt(500,1000))
	}
}