package main

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



//type person struct {
//	name string
//	city string
//	age  int8
//}
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

//func main()  {
//	var p2 = new(person)
//	p2.name = "测试"
//	p2.age = 18
//	p2.city = "北京"
//	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"测试", city:"北京", age:18}
//}
//使用&对结构体进行去地址操作相当于对该结构体类型进行了一次new实例化操作
//func main()  {
//	p3 := &person{}
//	fmt.Printf("%T\n", p3)     //*main.person
//	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
//	p3.name = "博客"
//	p3.age = 30
//	p3.city = "成都"
//	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
//}//p3.name = "博客"其实在底层是(*p3).name = "博客"，这是Go语言帮我们实现的语法糖。

//结构体初始化
//func main() {
//	var p4 person
//	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
//}


//func main() {
//	var p4 person
//	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
//	p5 := person{    //键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
//		name: "pprof.cn",
//		city: "北京",
//		age:  18,
//	}
//	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
//	p6 := &person{     //对结构体指针进行键值对初始化
//		name: "pprof.cn",
//		city: "北京",
//		age:  18,
//	}
//	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18}
//	p7 := &person{    //当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值
//		city: "北京",
//	}
//	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
//	p8 := &person{     //初始化结构体的时候可以简写，也就是初始化的时候不写键
//		"pprof.cn",
//		"北京",
//		18,
//	}
//	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
//}
//注意：
//1.必须初始化结构体的所有字段。
//2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
//3.该方式不能和键值初始化方式混用。


//###################结构体内存布局
//type test struct {
//	a int8
//	b int8
//	c int8
//	d int8
//}
//
//func main()  {
//	type test struct {
//		a int8
//		b int8
//		c int8
//		d int8
//	}
//	n := test{
//		1, 2, 3, 4,
//	}
//	fmt.Printf("n.a %p\n", &n.a)
//	fmt.Printf("n.b %p\n", &n.b)
//	fmt.Printf("n.c %p\n", &n.c)
//	fmt.Printf("n.d %p\n", &n.d)
//}
//输出结果：
//n.a 0xc000096004
//n.b 0xc000096005
//n.c 0xc000096006
//n.d 0xc000096007

//#######面试题
//type student struct {
//	name string
//	age  int
//}
//
//func main() {
//	m := make(map[string]*student)
//	stus := []student{
//		{name: "pprof.cn", age: 18},
//		{name: "测试", age: 23},
//		{name: "博客", age: 28},
//	}
//
//	for _, stu := range stus {
//		m[stu.name] = &stu   //for循环后m.name最后一个值是博客
//	}
//	for k, v := range m {   //m的值是map的stus
//		fmt.Println(k, "=>", v.name)
//	}
//}

//输出结果：
//pprof.cn => wolf
//测试 => wolf
//wolf => wolf








