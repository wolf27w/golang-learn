package main

import (
	"io"
	"log"
	"os"
)


func appendStringInFile(filePath, content string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("文件打开失败: %v", err)
	}
	defer file.Close()
	// 查找文件末尾的偏移量
	n, _ := file.Seek(2, io.SeekStart)
	// 从末尾的偏移量开始写入内容
	_, err = file.WriteAt([]byte("\n"+content), n)
	if err != nil {
		log.Fatalf("文件写入失败: %v", err)
	}
}

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/index.html"
	str1 := "hello www.wulaoer.org"
	appendStringInFile(path,str1)




}