package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reLineke = `&nbsp;&nbsp;&nbsp;&nbsp;[\u4e00-\u9fa5]`
)


// 处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func contente(url string)  {
	resp,err := http.Get(url)
	HandleError(err,"content error")
	defer resp.Body.Close()
	page,err := ioutil.ReadAll(resp.Body)
	HandleError(err,"iou")
	pages := string(page)
	re := regexp.MustCompile(reLineke)
	results := re.FindAllStringSubmatch(pages,-1)
	for _, resylt := range results {
		fmt.Println(resylt)
	}
}



func main() {
	path:="https://www.qbiqu.com/0_4/4235911.html"
	contente(path)
}