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
	//if err != nil {
	//	fmt.Printf("文件错误",err)
	//}
	//fmt.Printf(string(content))
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("文件错误")
	} else {
		//fmt.Printf("file=%v",&file)
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