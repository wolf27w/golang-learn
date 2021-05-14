//package main
//
//import "fmt"
//
//func main()  {
//	slice := []int{0,1,2,3}
//	m := make(map[int]*int)
//	for key,val := range slice {
//		value := val
//		m[key] = &value
//	}
//	for k,v := range m {
//		fmt.Println(k,"--->",*v)
//	}
//}

package main

//map中的数据都是成对出现的，语法结构为：map[KeyType]ValueType keyType表示键类型，valueType表示键对应的值的类型
//map类型的变量默认初始值为nil，需要使用make()函数累分配内存，语法为： make(map[KeyType]ValueType, [cap]  cap表示map的容量，改参数虽然不是必须的，但是在初始化map的时候就为其制定一个合适的容量
//func main() {
//	scoreMap := make(map[string]int, 8)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	fmt.Println(scoreMap)
//	fmt.Println(scoreMap["小明"])
//	fmt.Printf("type of a:%T\n", scoreMap)
//}
//输出结果：
//map[小明:100 张三:90]
//100
//type of a:map[string]int


///########################
//map支持在生命的时候填充元素

//func main() {
//	userInfo := map[string]string{
//		"username": "pprof.cn",
//		"password": "123456",
//	}
//	fmt.Println(userInfo) //map[password:123456 username:pprof.cn]
//}

//###############################
//判断map中键是否存在特殊的写法，格式：value, ok := map[key]

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
//	v, ok := scoreMap["张三"]
//	if ok {
//		fmt.Println(v) //输出结果90
//	} else {
//		fmt.Println("查无此人")
//	}
//}

//#############################
//使用rang遍历map

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["王五"] = 60
//	for k, v := range scoreMap {
//		fmt.Println(k, v)
//	}
//}
//输出结果：
//张三 90
//小明 100
//王五 60
//只想遍历key
//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["王五"] = 60
//	for k := range scoreMap {
//		fmt.Println(k)
//	}
//}
//输出结果：
//小明
//王五
//张三


//func main() {
//
//slice := []int{0,1,2,3}
//m := make(map[int]*int)
//fmt.Println()
//for key,val := range slice {
//	value := val
//	m[key] = &value
//}
//
//for k,v := range m {
//	fmt.Println(k,"===>",*v)
//}
//}