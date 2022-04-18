package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/README.md"
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("文件错误",err)
	}
	fmt.Printf("%v",content)

}