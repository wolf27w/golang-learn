package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	//fmt.Println(os.Getwd())
	path := "golang-learn/index.html"
	//path1 := "golang-learn/index1.html"
	srcfile,err := os.OpenFile(path,os.O_RDWR,0644)
	if err != nil {
		fmt.Println("open file filed %v",err)
		return
	}
	tmpfile,err := os.OpenFile("./123.tmp",os.O_CREATE | os.O_WRONLY | os.O_TRUNC,0644)
	if err != nil {
		fmt.Printf("open file filed %v",err)
		return
	}
	var b [10]byte //定义文件的需要插入的定位符
	n,err := srcfile.Read(b[:])
	if err != nil {
		fmt.Printf("read file failed,err:%v\n",err)
		return
	}
	tmpfile.Write(b[:n])
	tmpfile.WriteString("www.wulaoer.org")
	var x [1024]byte
	for {
		n,err := srcfile.Read(x[:])
		if err == io.EOF {
			tmpfile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n",err)
			return
		}
		tmpfile.Write(x[:n])
	}
	tmpfile.Close()
	srcfile.Close()
	ok := os.Rename("./123.tmp",path)
	if ok == nil {
		fmt.Printf(" sucnss")
	} else {
		fmt.Printf("error")
	}

}