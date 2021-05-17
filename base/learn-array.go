package main

import "fmt"

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





//下面这个函数访问越界，如飞下标在数组合法范围内，则触发访问越界，会panic。数组是值类型，赋值和传参会复制整个数组，因此，改变副本的值，不会改变本身的值，对"=="和"!="操作符，因为内存总是呗初始化过的，指针数组[n]*T，数组指针*[n]T
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


//##初始化数组，针对不同类型的数组
//全局变量
//var arr0 [2][3]int = [...][3]int{{1,2,3},{7,8,9}}
//var arr0 [5]int = [5]int{1,2,3}    //[1 2 3 0 0]
//var arr1 = [5]int{1,2,3,4,5}       //[1 2 3 4 5]
//var arr2 = [...]int{1,2,3,4,5,6}   //[1 2 3 4 5 6]
//var str = [5]string{3: "hello word",4:"Tom"}  //[   hello word Tom]
////局部变量
//func main()  {
//	a := [3]int{1,2}     //[1 2 0]
//	b := [...]int{1,2,3,4}   //[1 2 3 4]
//	c := [5]int{2:100,4:200}  //[0 0 100 0 200]
//	d := [...]struct {        //[{user1 10} {user2 20}]
//		name string
//		age uint8
//	}{
//		{"user1",10},
//		{"user2",20},
//	}
//	fmt.Println(arr0,arr1,arr2,str)
//	fmt.Println(a,b,c,d)
//}
//    全局
//    var arr0 [5][3]int   //
//    var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
//    局部：
//    a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
//    b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."【长度不能使用...】。


//var arr0 [5][3]int   //[[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]
//var arr1 [2][3]int = [...][3]int{{1,2,3},{7,8,9}}  //[[1 2 3] [7 8 9]]
//
//func main()  {
//	a := [2][3]int{{1,2,3},{4,5,6}}  //[[1 2 3] [4 5 6]]
//	b := [...][2]int{{1,1},{2,2},{3,3}}  //[[1 1] [2 2] [3 3]]
//	fmt.Println(arr0,arr1)
//	fmt.Println(a,b)
//}
//值拷贝
//func text(x [2]int)  {
//	fmt.Printf("x: %p\n", &x)
//	x[1] = 1000
//}
//func main()  {
//	a := [2]int{}
//	fmt.Printf("a: %p\n",&a)  //使用了printf内存地址
//	text(a)
//	fmt.Println(a)
//}

//内置函数len和cap都繁华数组长度(元素数量)

//func main()  {
//	a := [2]int{}
//	print(len(a),cap(a))
//}

//多维数组变量
//
//func main()  {
//	var f [2][3]int =[...][3]int{{1,2,3},{7,8,9}}
//	for k1,v1 := range f {
//		for k2,v2 := range v1 {
//			fmt.Printf("(%d,%d)=%d",k1,k2,v2)
//		}
//		fmt.Println()
//	}
//}
//输出结果：
//(0,0)=1(0,1)=2(0,2)=3
//(1,0)=7(1,1)=8(1,2)=9

//数组拷贝和传参

//func test(arr *[5]int)  {
//	arr[0] = 10
//	for i,v := range arr {
//		fmt.Println(i,v)
//	}
//}
//func main()  {
//	var arr1 [5]int
//	test(&arr1)
//	fmt.Println(arr1)
//	arr2 := [...]int{2,3,4,8,10}
//	test(&arr2)
//	fmt.Println(arr2)
//}
//输出结果：
//0 10
//1 0
//2 0
//3 0
//4 0
//[10 0 0 0 0]
//0 10
//1 3
//2 4
//3 8
//4 10
//[10 3 4 8 10]

//数组练习，求数组所有元素只和

//func sumarr(a [10]int) int  {//根据下标统计元素只和
//	var sum int = 0
//	for i := 0; i < len(a); i++ {
//		sum += a[i]
//	}
//	return sum
//}
//
//func main()  {
//	rand.Seed(time.Now().Unix())
//	var b [10]int
//
//	for i := 0; i < len(b);i++ {  //随机生成0到1000的随机数，赋值给b
//		b[i] = rand.Intn(1000)
//	}
//	fmt.Println(b)
//	sum := sumarr(b)
//	fmt.Printf("sum=%d\n",sum)
//}

//求两个元素只和等于8的下标
func myTest(a [5]int, target int) {
	// 遍历数组
	for i := 0; i < len(a); i++ {
		other := target - a[i]
		// 继续遍历
		for j := i + 1; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
//输出结果：
//(0,4)
//(1,2)


