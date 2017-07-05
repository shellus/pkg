package rand

import (
	"crypto/rand"
	"math/big"
)

func Int64(min, max int64) int64 {

	i, err := rand.Int(rand.Reader, new(big.Int).SetInt64(max - min))
	if err != nil {
		panic(err)
	}
	return i.Int64() + min
}
func Int(min, max int) int {
	return int(Int64(int64(min), int64(max)))
}