package main

//################面向对象

//go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//    go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//人
//type Person struct {
//	name string
//	sex  string
//	age  int
//}
//
//type Student struct {
//	Person
//	id   int
//	addr string
//}
//
//func main() {
//	// 初始化
//	s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
//	fmt.Println(s1)
//
//	s2 := Student{Person: Person{"5lmh", "man", 20}}
//	fmt.Println(s2)
//
//	s3 := Student{Person: Person{name: "5lmh"}}
//	fmt.Println(s3)
//}
//输出结果

//{{5lmh man 20} 1 bj}
//{{5lmh man 20} 0 }
//{{5lmh  0} 0 }

//同名字段的情况

//人
//type Person struct {
//	name string
//	sex  string
//	age  int
//}
//
//type Student struct {
//	Person
//	id   int
//	addr string
//	//同名字段
//	name string
//}
//
//func main() {
//	var s Student
//	// 给自己字段赋值了
//	s.name = "5lmh"
//	fmt.Println(s)
//
//	// 若给父类同名字段赋值，如下
//	s.Person.name = "枯藤"
//	fmt.Println(s)
//}
//输出结果
//{{  0} 0  5lmh}
//{{枯藤  0} 0  5lmh}

//所有的内置类型和自定义类型都是可以作为匿名字段去使用


//人
//type Person struct {
//	name string
//	sex  string
//	age  int
//}
//
//// 自定义类型
//type mystr string
//
//// 学生
//type Student struct {
//	Person
//	int
//	mystr
//}
//
//func main() {
//	s1 := Student{Person{"5lmh", "man", 18}, 1, "bj"}
//	fmt.Println(s1)
//}

//输出结果
//{{5lmh man 18} 1 bj}


//指针类型匿名字段

//人
//type Person struct {
//	name string
//	sex  string
//	age  int
//}
//
//// 学生
//type Student struct {
//	*Person
//	id   int
//	addr string
//}
//
//func main() {
//	s1 := Student{&Person{"5lmh", "man", 18}, 1, "bj"}
//	fmt.Println(s1)
//	fmt.Println(s1.name)
//	fmt.Println(s1.Person.name)
//}
//输出结果
//{0xc000074180 1 bj}
//5lmh
//5lmh

//########################接口

//接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。

//interface是一组method的集合，是duck-type programming的一种体现。接口做的事情就像是定义一个协议（规则），只要一台机器有洗衣服和甩干的功能，我就称它为洗衣机。不关心属性（数据），只关心行为（方法）

//type Cat struct{}
//
//func (c Cat) Say() string { return "喵喵喵" }
//
//type Dog struct{}
//
//func (d Dog) Say() string { return "汪汪汪" }
//
//func main() {
//    c := Cat{}
//    fmt.Println("猫:", c.Say())
//    d := Dog{}
//    fmt.Println("狗:", d.Say())
//}

//上面的代码中定义了猫和狗，然后它们都会叫，你会发现main函数中明显有重复的代码，如果我们后续再加上猪、青蛙等动物的话，我们的代码还会一直重复下去。那我们能不能把它们当成“能叫的动物”来处理呢？
//
//像类似的例子在我们编程过程中会经常遇到：


//像类似的例子在我们编程过程中会经常遇到：
//
//比如一个网上商城可能使用支付宝、微信、银联等方式去在线支付，我们能不能把它们当成“支付方式”来处理呢？
//
//比如三角形，四边形，圆形都能计算周长和面积，我们能不能把它们当成“图形”来处理呢？
//
//比如销售、行政、程序员都能计算月薪，我们能不能把他们当成“员工”来处理呢？
//
//Go语言中为了解决类似上面的问题，就设计了接口这个概念。接口区别于我们之前所有的具体类型，接口是一种抽象的类型。当你看到一个接口类型的值时，你不知道它是什么，唯一知道的是通过它的方法能做什么。

//#################接口的定义

//接口是一个或多个方法签名的集合。
//    任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
//    就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
//    这称为Structural Typing。
//    所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
//    当然，该类型还可以有其他方法。
//
//    接口只有方法声明，没有实现，没有数据字段。
//    接口可以匿名嵌入其他接口，或嵌入到结构中。
//    对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
//    只有当接口存储的类型和对象都为nil时，接口才等于nil。
//    接口调用不会做receiver的自动转换。
//    接口同样支持匿名字段方法。
//    接口也可实现类似OOP中的多态。
//    空接口可以作为任何类型数据的容器。
//    一个类型可实现多个接口。
//    接口命名习惯以 er 结尾。


//############每个接口由数个方法组成，接口的定义格式如下：


// type 接口类型名 interface{
//        方法名1( 参数列表1 ) 返回值列表1
//        方法名2( 参数列表2 ) 返回值列表2
//        …
//    }
