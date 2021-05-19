package main

import "fmt"

//golang没有"类"的概念，也不支持"类"的继承等面向对象的概念，golang中通过结构体的内嵌在配合接口比面向对象具有更高的扩展性和灵活性。
//golan中，基本的数据类型，如string，整型，浮点型，布尔等数据类型，可以使用type关键字来定义类型
//将MyInt定义为int类型
//type MyInt int
//除了int的类型，还有byte，rune等类型，类型别名与类型定义表面上看是只有一个等号等差异，但是下面看看：

//类型定义
//type NewInt int
////类型别名
//type MyInt = int
//func main() {
//	var a NewInt
//	var b MyInt
//	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt  //结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型
//	fmt.Printf("type of b:%T\n", b) //type of b:int     //b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
//}
//输出结果：
//type of a:main.NewInt
//type of b:int

//###########################结构体
//结构体是一种自定义数据类型，可以封装多个基本数据类型。英文名称struct，也就可以通过struct来定义自己的类型了。
//语法：
//type 类型名 struct {
//	字段名 字段类型
//	字段名 字段类型
//	…
//}
//1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
//2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
//3.字段类型：表示结构体字段的具体类型。
//举例：
//type person struct {
//	name string
//	city string
//	age  int8
//}
//或者
//type person1 struct {
//	name, city string
//	age        int8
//}
//这样我们就拥有了一个person的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄



type person struct {
	name string
	city string
	age  int8
}
//func main() {
//	var p1 person
//	p1.name = "pprof.cn"    //通过.来访问结构体的字段（成员变量）
//	p1.city = "北京"
//	p1.age = 18
//	fmt.Printf("p1=%v\n", p1)  //p1={pprof.cn 北京 18}
//	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"pprof.cn", city:"北京", age:18}
//}


//匿名结构体

//func main() {
//	var user struct{Name string; Age int}
//	user.Name = "pprof.cn"
//	user.Age = 18
//	fmt.Printf("%#v\n", user) //struct { Name string; Age int }{Name:"pprof.cn", Age:18}
//}
//通过new关键字对结构体进行实例化，得到结构体对地址，语法如下：
//func main()  {
//	var p2 = new(person)
//	fmt.Printf("%T\n", p2)     //*main.person
//	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
//}

func main()  {
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
}





