package sutil

import (
	"os/user"
	"os"
	"math/big"
	"crypto/rand"
)

func HomeDir() string {
	usr, err := user.Current()
	var homeDir string
	if err == nil {
		homeDir = usr.HomeDir
	} else {
		// Maybe it's cross compilation without cgo support. (darwin, unix)
		homeDir = os.Getenv("HOME")
	}
	return homeDir
}

func RandInt64(min, max int64) int64 {

	i, err := rand.Int(rand.Reader, new(big.Int).SetInt64(max - min))
	if err != nil {
		panic(err)
	}
	return i.Int64() + min
}
func RandInt(min, max int) int {
	return int(RandInt64(int64(min), int64(max)))
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return true
}