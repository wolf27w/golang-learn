package main

import (
	"fmt"
	"io/ioutil"
	"mahonia"
	"net/http"
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
	var enc mahonia.Decoder

	enc = mahonia.NewDecoder("gbk")
	resp,err := http.Get("https://www.qbiqu.com/7_7351/")
	HandleError(err,"http.Get url")
	defer resp.Body.Close()
	pageByte,err := ioutil.ReadAll(resp.Body)
	HandleError(err,"ioutil,ReadAll")
	pagestr := string(pageByte)
	fmt.Println("UTF-8 to GBK:", enc.ConvertString(pagestr))
	//re := regexp.MustCompile(reQQEmail)
	//results := re.FindAllStringSubmatch(pagestr,-1)
	//for _, result := range results {
	//	fmt.Println("email:",result[0])
	//	fmt.Println("qq:",result[1])
	//}
}



func main()  {
	novel()
	
}