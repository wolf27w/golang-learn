package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	//fmt.Println(os.Getwd())
	path := "golang-learn/index.html"
	//content, err := os.OpenFile(path,os.O_WRONLY | os.O_CREATE, 0655)
	//if err != nil {
	//	fmt.Println("文件错误",err)
	//	return
	//}
	//defer content.Close()
	//str := "hello golang \r\n"
	//write := bufio.NewWriter(content)
	////for i := 0; i < 5; i++ {
	////	write.WriteString(str)
	////}
	//write.WriteString(str)
	//write.Flush()//刷新缓存写入到文件中

	//con, err := os.Open(path)
	//if err != nil {
	//	fmt.Print("文件读取错误")
	//}
	//defer con.Close()
	//reader := bufio.NewReader(con)
	//for {
	//	strs,err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Print(err)
	//		break
	//	}
	//	fmt.Print(strs)
	//}
	//fmt.Println("文件结束")


	file, err := os.Open(path)
	if err != nil {
		fmt.Println("打开文件错误")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		str = strings.TrimSpace(str)
		fmt.Println(str)
		if err == io.EOF {
			break
		}
		//fmt.Print(str)
	}
	//fmt.Println("文件结束")

	//file , err := ioutil.ReadFile(path)
	//if err != nil {
	//	fmt.Printf("read fiel err = %v",err)
	//}
	//fmt.Printf(string(file))
}