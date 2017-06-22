package proxy_pool_test

import (
	"testing"
	"github.com/shellus/pkg/proxy_pool"
)

func TestExample(t *testing.T) {
	c := proxy_pool.Client{}
	c.NextProxy()
	c.Get("http://httpbin.org/ip")
}