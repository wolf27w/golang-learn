package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// \d是数字
	reQQEmail = `(\d+)@qq.com`
)

func HandleError(err error,why string)  {
	if err != nil {
		fmt.Printf(why,err)
	}
}

func novel()  {

	resp,err := http.Get("https://www.qbiqu.com/7_7351/14281108.html")
	HandleError(err,"http.Get url")
	defer resp.Body.Close()
	pageByte,err := ioutil.ReadAll(resp.Body)
	HandleError(err,"ioutil,ReadAll")
	pagestr := string(pageByte)
	//fmt.Println(pagestr)
	re := regexp.MustCompile()
	results := re.FindAllStringSubmatch(pagestr,-1)
	fmt.Println(results)
	//for _, result := range results {
	//	fmt.Println("email:",result[0])
	//	fmt.Println("qq:",result[1])
	//}
}



func main()  {
	novel()
	
}