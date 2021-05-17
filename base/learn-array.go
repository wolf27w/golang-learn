package main

//import "fmt"

//数组(array)：是同一种数据类型元素的集合。申明确定，使用修改，但是数组大小不可更改,语法：var 数组变量名 [元素数量]T
//var a [2]int
//数组定义：var a [len]int, 比如：var a [5]int,数组长度必须是常量，切类型是组成部分，一旦定义，长度不变，长度是数组类型的一部分，因此，var a[5] int和var a[10]int 是不同类型。数组可以通过下标进行访问，下标是从0开始，最后一个是len-1
//####################数组初始化
//func main()  {
//	var wolf27  [3]int
//	var wulaoer = [3]int{1,2}
//	var city = [3]string{"北京","上海","深圳"}
//	fmt.Println(wolf27)   //[0 0 0]
//	fmt.Println(wulaoer)  //[1 2 0]
//	fmt.Println(city)   //[北京 上海 深圳]
//}
//数组初始化，可以使用初始化列表来设置初始化元素的值，







//func main()  {
//	var a[5] int
//	var b[10] int
//	for i := 0; i < len(b); i ++{ //for遍历
//		fmt.Println(i)
//	}
//	for index,k := range a{ //for...range遍历
//		fmt.Println(index,k)
//	}
//}
//输出结果：
//0
//1
//2
//3
//4
//5
//6
//7
//8
//9
//0 0
//1 0
//2 0
//3 0
//4 0

//func main() {
//	var a = [...]string{"北京", "上海", "深圳"}
//	// 方法1：for循环遍历
//	for i := 0; i < len(a); i++ {
//		fmt.Println(a[i])
//	}
//
//	// 方法2：for range遍历
//	for index, value := range a {
//		fmt.Println(index, value)
//	}
//}