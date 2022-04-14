package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	//file , err := os.Open("golang-learn/go.mod")
	//如果是较小文件可以使用
	file := "golang-learn/go.mod"
	//content, err := ioutil.ReadFile(file)
	content, err := os.OpenFile(file,os.O_WRONLY | os.O_CREATE, 077)
	if err != nil {
		fmt.Println("file错误信息",err)
		return
	}

	defer content.Close()
	str := "wulaoer.org\r\n"
	writer := bufio.NewWriter(content)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
