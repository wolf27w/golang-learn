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

//###################构造函数
//原本golang结构体里没有构造函数，下面我们自己实现一个person的结构函数，因为struct是值类型，

//func newPerson(name, city string, age int8) *person {
//	return &person{
//		name: name,
//		city: city,
//		age:  age,
//	}
//}
//func main()  {
//	p9 := newPerson("pprof.cn", "测试", 90)
//	fmt.Printf("%#v\n", p9)  //&main.person{name:"pprof.cn", city:"测试", age:90}
//}


//方法和接收者
//Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
//语法：
//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	函数体
//}

//1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
//2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
//3.方法名、参数列表、返回参数：具体格式与函数定义相同。


//Person 结构体
//type Person struct {
//	name string
//	age  int8
//}
////NewPerson 构造函数
//func NewPerson(name string, age int8) *Person {
//	return &Person{
//		name: name,
//		age:  age,
//	}
//}
////Dream Person做梦的方法
//func (p Person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name) //测试的梦想是学好Go语言！
//}
//
//func main() {
//	p1 := NewPerson("测试", 25)
//	p1.Dream()
//}
//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。


//指针类型接受者
//指针类型接受者是有一个结构体的指针组成，由于指针的特效，调用方法时修改指针的任意成员变量，在方法结束后，修改都是有效的，这种方式就十分接近于其他语言中面向对象的this或者self，，例如：

//type Person struct {
//	name string
//	age  int8
//}
//
////NewPerson 构造函数
//func NewPerson(name string, age int8) *Person {
//	return &Person{
//		name: name,
//		age:  age,
//	}
//}
//////Dream Person做梦的方法
//func (p Person) Dream() {
//	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
//}
//////SetAge2 Person做梦的方法
//func (p Person) SetAge2(newAge int8) {
//	p.age = newAge
//}
//func main() {
//	p1 := NewPerson("测试", 25)
//	p1.Dream()
//	fmt.Println(p1.age) // 25
//	p1.SetAge2(30) // (*p1).SetAge2(30)
//	fmt.Println(p1.age) // 25
//}
//Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

//什么时候应该使用指针类型接收者
//1.需要修改接收者中的值
//2.接收者是拷贝代价比较大的大对象
//3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

//任意类型添加方法，接受者类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法，
//举例：基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

//MyInt 将int定义为自定义MyInt类型
//type MyInt int
//
////SayHello 为MyInt添加一个SayHello的方法
//func (m MyInt) SayHello() {
//	fmt.Println("Hello, 我是一个int。")
//}
//func main() {
//	var m1 MyInt
//	m1.SayHello() //Hello, 我是一个int。
//	m1 = 100
//	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
//}
//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。


//######结构体匿名字段
//结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
//Person 结构体Person类型
//type Person struct {
//	string
//	int
//}
//func main() {
//	p1 := Person{
//		"pprof.cn",
//		18,
//	}
//	fmt.Printf("%#v\n", p1)        //main.Person{string:"pprof.cn", int:18}
//	fmt.Println(p1.string, p1.int) //pprof.cn 18
//}

//匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。


//#######################嵌套结构体
//一个结构体中可以嵌套包含另一个结构体或结构体指针。
//Address 地址结构体
//type Address struct {
//	Province string
//	City     string
//}
////User 用户结构体
//type User struct {
//	Name    string
//	Gender  string
//	Address Address
//}
//func main() {
//	user1 := User{
//		Name:   "pprof",
//		Gender: "女",
//		Address: Address{
//			Province: "黑龙江",
//			City:     "哈尔滨",
//		},
//	}
//	fmt.Printf("user1=%#v\n", user1)//user1=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
//}


//############## 嵌套匿名结构体

////Address 地址结构体
//type Address struct {
//	Province string
//	City     string
//}
////User 用户结构体
//type User struct {
//	Name    string
//	Gender  string
//	Address //匿名结构体
//}
//func main() {
//	var user2 User
//	user2.Name = "pprof"
//	user2.Gender = "女"
//	user2.Address.Province = "黑龙江"    //通过匿名结构体.字段名访问
//	user2.City = "哈尔滨"                //直接访问匿名结构体的字段名
//	fmt.Printf("user2=%#v\n", user2) //user2=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"黑龙江", City:"哈尔滨"}}
//}

//当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

//#############嵌套结构体的字段名冲突

//嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。

//Address 地址结构体
//type Address struct {
//	Province   string
//	City       string
//	CreateTime string
//}
////Email 邮箱结构体
//type Email struct {
//	Account    string
//	CreateTime string
//}
////User 用户结构体
//type User struct {
//	Name   string
//	Gender string
//	Address
//	Email
//}
//func main() {
//	var user3 User
//	user3.Name = "pprof"
//	user3.Gender = "女"
//	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
//	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
//	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
//	fmt.Printf("user3=%#v\n",user3)//ser3=main.User{Name:"pprof", Gender:"女", Address:main.Address{Province:"", City:"", CreateTime:"2000"}, Email:main.Email{Account:"", CreateTime:"2000"}}
//}








