package proxy_pool_test

import (
	"testing"
	//"github.com/shellus/pkg/proxy_pool"
	//"net/http/httptrace"
	//"fmt"
)

func TestExample(t *testing.T) {
	//c := proxy_pool.Client{}
	//c.NextProxy()
	//c.Get("http://httpbin.org/ip")
	//
	//
	//req, _ := http.NewRequest("GET", "https://httpbin.org/ip", nil)
	//trace := &httptrace.ClientTrace{
	//	GotConn: func(connInfo httptrace.GotConnInfo) {
	//		// todo 在这里入选代理IP
	//		fmt.Printf("Got Conn: %+v\n", connInfo)
	//	},
	//	DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
	//		//fmt.Printf("DNS Info: %+v\n", dnsInfo)
	//	},
	//}
	//req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	//_, err := http.DefaultClient.Do(req)
	//
	//if err != nil {
	//	fmt.Printf("%#v", err)
	//}
}