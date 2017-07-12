package crawl

import (
	"context"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Crawl struct {

}
type Page struct {
	Url url.URL
	resp http.Response
	goquery *goquery.Document
}
type Route struct {
	rule string
	tunnel chan *Page
}

func NewCrawler() *Crawl{
	return &Crawl{}
}

// 加入url
func (craw *Crawl) JoinUrl(ctx context.Context){

}
// 开始抓取
func (craw *Crawl) Work(ctx context.Context){

}

// 监听指定规则过滤后的页面
func (craw *Crawl) Filter(rule string)(chan *Page){
	r := new(Route)
	r.tunnel = make(chan *Page)
	r.rule = rule
	return r.tunnel
}

// 获取路由参数
func (craw *Page) Param(name string)(string, error){

}

func (craw *Page) GoQuery()(goquery.Selection){
	var err error

	if craw.goquery == nil {
		craw.goquery, err = goquery.NewDocumentFromReader(craw.resp.Body)
		if err != nil {
			panic(err)
		}
	}

	return craw.goquery
}
