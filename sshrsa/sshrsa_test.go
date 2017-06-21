package sshrsa

import (
	"testing"
	"errors"
)

func TestCrypt(t *testing.T) {
	str := "123123123123123123123123123123123123123123123"
	c, err := Encrypt([]byte(str))
	if err != nil {
		t.Error(err)
	}
	b, err := Decrypt(c)
	if err != nil {
		t.Error(err)
	}
	if string(b) != str {
		t.Error(errors.New("err"))
	}
}