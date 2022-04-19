package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/index5.html"
	//content, err := os.OpenFile(path,os.O_WRONLY | os.O_TRUNC, 0655)
	content, err := os.OpenFile(path,os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0655)
	if err != nil {
		fmt.Println("文件错误",err)
		return
	}
	defer content.Close()
	str := "hello golang \r\n"
	write := bufio.NewWriter(content)
	write.WriteString(str)
	write.Flush()//刷新缓存写入到文件中
}