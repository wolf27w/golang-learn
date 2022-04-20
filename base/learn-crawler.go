package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"), colly.MaxDepth(1), colly.Debugger(&debug.LogDebugger{}))
	//文章列表
	//c.OnHTML("div[class='dahengfu']", func(e *colly.HTMLElement) {
	//	//列表中每一项
	//	e.ForEach("div", func(i int, item *colly.HTMLElement) {
	//		//文章链接
	//		href := item.ChildAttr("a[href='/'] > a[class='title']", "href")
	//		//文章标题
	//		title := item.ChildText("div[class='bookname'] > a[class='title']")
	//		//文章摘要
	//		summary := item.ChildText("div[id='content'] > p[class='abstract']")
	//		fmt.Println(title, href)
	//		fmt.Println(summary)
	//		//fmt.Println()
	//	})
	//})
	c.OnHTML("div[class='box_con']", func(e *colly.HTMLELement) {
		e.FOREACH("li", func(i int, item *colly.HTMLElement) {
			href := item.ChildAttr("div[class='nav']" > a[class='header'],)
		})
	})
	err := c.Visit("https://www.qbiqu.com/7_7351/14281109.html")
	if err != nil {
		fmt.Println(err.Error())
	}
}





//func main() {
//	path:="https://www.qbiqu.com/0_4/4235911.html"
//	contente(path)
//}