package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/index.html"
	file,err := os.Open(path)
	if err != nil {
		fmt.Println("file filed err %v",err)
	}
	defer file.Close()

	//content, err := os.OpenFile(path,os.O_WRONLY | os.O_TRUNC, 0655)
	////content, err := os.OpenFile(path,os.O_RDWR | os.O_TRUNC | os.O_CREATE, 0655)//两者都有覆盖的意思
	//if err != nil {
	//	fmt.Println("文件错误",err)
	//	return
	//}
	//defer content.Close()
	//str1 := "hello www.wulaoer.org"
	read := bufio.NewReader(file)
	fmt.Print(read)




}