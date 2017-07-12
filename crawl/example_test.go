package crawl_test

import (
	"testing"
	"context"
	"os/signal"
	"os"
	"fmt"
	"github.com/shellus/pkg/crawl"
)

func TestExample(t *testing.T) {


	crawler := crawl.NewCrawler()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go crawler.Work(ctx)


	for page := range crawler.Filter("/article/{id}.html") {
		articleName, _ := page.Param("id")
		articleTitle := page.GoQuery().Find("title").Text()
		articleContent, err := page.GoQuery().Find(".article-content").Html()
		if err != nil {
			panic(err)
		}

		fmt.Println(articleName, articleTitle, articleContent)

		// 再次加入
		crawler.JoinUrl(page.Url)
	}

	crawler.JoinUrl("http://blog.endaosi.com/")

	ListenCtrlC()
}

func ListenCtrlC(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	s := <-c

	// 这里会返回去执行cancel(),所以就不用管啦
	fmt.Println("Got signal:", s)
}