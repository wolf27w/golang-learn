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


//其中

//1.接口名：使用type将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加er，如有写操作的接口叫Writer，有字符串功能的接口叫Stringer等。接口名最好要能突出该接口的类型含义。
//    2.方法名：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
//    3.参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以省略。

//type writer interface{
//    Write([]byte) error
//}

//当你看到这个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的Write方法来做一些事情。


//#################实现接口的条件

//一个对象只要全部实现了接口中的方法，那么就实现了这个接口。换句话说，接口就是一个需要实现的方法列表。

//我们来定义一个Sayer接口：

//// Sayer 接口
//type Sayer interface {
//    say()
//}

//定义dog和cat两个结构体：

//type dog struct {}
//
//type cat struct {}

//因为Sayer接口里只有一个say方法，所以我们只需要给dog和cat 分别实现say方法就可以实现Sayer接口了。

//// dog实现了Sayer接口
//func (d dog) say() {
//    fmt.Println("汪汪汪")
//}
//
//// cat实现了Sayer接口
//func (c cat) say() {
//    fmt.Println("喵喵喵")
//}

//接口的实现就是这么简单，只要实现了接口中的所有方法，就实现了这个接口。

//###################接口类型######################

//接口类型变量能够存储所有实现了该接口的实例。 例如上面的示例中，Sayer类型的变量能够存储dog和cat类型的变量。

//func main() {
//	var x Sayer // 声明一个Sayer类型的变量x
//	a := cat{}  // 实例化一个cat
//	b := dog{}  // 实例化一个dog
//	x = a       // 可以把cat实例直接赋值给x
//	x.say()     // 喵喵喵
//	x = b       // 可以把dog实例直接赋值给x
//	x.say()     // 汪汪汪
//}

//值接收者和指针接收者实现接口的区别

//我们有一个Mover接口和一个dog结构体。

//type Mover interface {
//    move()
//}
//
//type dog struct {}

//值接收者实现接口

//func (d dog) move() {
//    fmt.Println("狗会动")
//}

//此时实现接口的是dog类型：

//func main() {
//    var x Mover
//    var wangcai = dog{} // 旺财是dog类型
//    x = wangcai         // x可以接收dog类型
//    var fugui = &dog{}  // 富贵是*dog类型
//    x = fugui           // x可以接收*dog类型
//    x.move()
//}

//使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui


//#########################指针接收者实现接口

//func (d *dog) move() {
//    fmt.Println("狗会动")
//}
//func main() {
//    var x Mover
//    var wangcai = dog{} // 旺财是dog类型
//    x = wangcai         // x不可以接收dog类型
//    var fugui = &dog{}  // 富贵是*dog类型
//    x = fugui           // x可以接收*dog类型
//}

//此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值

//#####################下面的代码是一个比较好的面试题

//请问下面的代码是否能通过编译？

//type People interface {
//	Speak(string) string
//}
//
//type Student struct{}
//
//func (stu *Stduent) Speak(think string) (talk string) {
//	if think == "sb" {
//		talk = "你是个大帅比"
//	} else {
//		talk = "您好"
//	}
//	return
//}
//
//func main() {
//	var peo People = Student{}
//	think := "bitch"
//	fmt.Println(peo.Speak(think))
//}


//######################类型与接口的关系

//####################一个类型实现多个接口

//一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。

//// Sayer 接口
//type Sayer interface {
//    say()
//}
//
//// Mover 接口
//type Mover interface {
//    move()
//}

//dog既可以实现Sayer接口，也可以实现Mover接口。

//type dog struct {
//    name string
//}
//
//// 实现Sayer接口
//func (d dog) say() {
//    fmt.Printf("%s会叫汪汪汪\n", d.name)
//}
//
//// 实现Mover接口
//func (d dog) move() {
//    fmt.Printf("%s会动\n", d.name)
//}
//
//func main() {
//    var x Sayer
//    var y Mover
//
//    var a = dog{name: "旺财"}
//    x = a
//    y = a
//    x.say()
//    y.move()
//}


//##############多个类型实现同一接口

//Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须由一个move方法。

//// Mover 接口
//type Mover interface {
//    move()
//}

//例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

//type dog struct {
//    name string
//}
//
//type car struct {
//    brand string
//}
//
//// dog类型实现Mover接口
//func (d dog) move() {
//    fmt.Printf("%s会跑\n", d.name)
//}
//
//// car类型实现Mover接口
//func (c car) move() {
//    fmt.Printf("%s速度70迈\n", c.brand)
//}

//这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的move方法就可以了



//func main() {
//	var x Mover
//	var a = dog{name: "旺财"}
//	var b = car{brand: "保时捷"}
//	x = a
//	x.move()
//	x = b
//	x.move()
//}

//输出结果
//    旺财会跑
//    保时捷速度70迈

//并且一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

//// WashingMachine 洗衣机
//type WashingMachine interface {
//    wash()
//    dry()
//}
//
//// 甩干器
//type dryer struct{}
//
//// 实现WashingMachine接口的dry()方法
//func (d dryer) dry() {
//    fmt.Println("甩一甩")
//}
//
//// 海尔洗衣机
//type haier struct {
//    dryer //嵌入甩干器
//}
//
//// 实现WashingMachine接口的wash()方法
//func (h haier) wash() {
//    fmt.Println("洗刷刷")
//}

















