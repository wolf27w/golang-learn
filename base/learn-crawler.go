package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func listAll(path string,curLayer int)  {
	//path := "/Users/wolf/Documents/GitHub"
	fileinfoos,err := ioutil.ReadDir(path)
	if err!= nil {
		fmt.Println("dir file filed %v",err)
		return
	}
	for _, info := range fileinfoos {
		if info.IsDir() {
			for temHier := curLayer;temHier>0;temHier--{
				fmt.Println("|\t")
			}
			fmt.Println(info.Name(),"\\")
			listAll(path+"/"+info.Name(),curLayer+1)
		} else {
			for temHier:= curLayer;temHier>0;temHier--{
				fmt.Println("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}

func main()  {
	dir := os.Args
	//fmt.Println(os.Getwd())
	listAll(dir,0)
}