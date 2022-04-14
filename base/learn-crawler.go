package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {

	file , err := os.Open("golang-learn/base/learn-if.go")
	if err != nil {
		fmt.Println("file错误信息",err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件结束")
}
