package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/README.md"
	//content, err := ioutil.ReadFile(path)
	////content , err := os.Open(path)
	//if err != nil {
	//	fmt.Println("文件错误",err)
	//}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("文件错误")
	}
	defer file.Close()
	read := bufio.NewReader(file)
	for {
		str,err := read.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

}