package main
//
//import (
//	"fmt"
//	"net/http"
//)

//###########################互联网协议介绍

//互联网的核心是一系列协议，总称为”互联网协议”（Internet Protocol Suite），正是这一些协议规定了电脑如何连接和组网。我们理解了这些协议，就理解了互联网的原理。由于这些协议太过庞大和复杂，没有办法在这里一概而全，只能介绍一下我们日常开发中接触较多的几个协议。

//################互联网分层模型

//互联网的逻辑实现被分为好几层。每一层都有自己的功能，就像建筑物一样，每一层都靠下一层支持。用户接触到的只是最上面的那一层，根本不会感觉到下面的几层。要理解互联网就需要自下而上理解每一层的实现的功能。



//package main
//
//import (
//"fmt"
//"net/http"
//)
//
//func main() {
//	//http://127.0.0.1:8000/go
//	// 单独写回调函数
//	http.HandleFunc("/go", myHandler)
//	//http.HandleFunc("/ungo",myHandler2 )
//	// addr：监听的地址
//	// handler：回调函数
//	http.ListenAndServe("127.0.0.1:8000", nil)
//}
//
//// handler函数
//func myHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println(r.RemoteAddr, "连接成功")
//	// 请求方式：GET POST DELETE PUT UPDATE
//	fmt.Println("method:", r.Method)
//	// /go
//	fmt.Println("url:", r.URL.Path)
//	fmt.Println("header:", r.Header)
//	fmt.Println("body:", r.Body)
//	// 回复
//	w.Write([]byte("www.5lmh.com"))
//}


