package sutil_test

import (
	"testing"
	"github.com/shellus/pkg/sutil"
	"path/filepath"
)

func TestHomeDir(t *testing.T) {
	t.Log(filepath.Join(sutil.HomeDir(),".ssh","id_rsa.pub"))

}