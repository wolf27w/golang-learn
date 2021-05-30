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








