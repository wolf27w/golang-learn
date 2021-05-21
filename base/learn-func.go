package main

//函数的特点
//• 无需声明原型。
//• 支持不定 变参。
//• 支持多返回值。
//• 支持命名返回参数。
//• 支持匿名函数和闭包。
//• 函数也是一种类型，一个函数可以赋值给变量。
//• 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
//• 不支持 重载 (overload)
//• 不支持 默认参数 (default parameter)。

//############函数申明
//函数声明包含一个函数名，参数列表， 返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。
//函数可以没有参数或接受多个参数。
//注意类型在变量名之后 。
//当两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。
//函数可以返回任意数量的返回值。
//使用关键字 func 定义函数，左大括号依旧不能另起一行。

//func sway(x, y int, s string) (int, string) {
//	// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
//	n := x + y
//	return n, fmt.Sprintf("你好%s,age%d",s, n)
//}
//func main()  {
//	var a,b int = 10,20
//	var s string ="wulaoer.org"
//	fmt.Println(sway(a,b,s))
//}
//输出结果：30 你好wulaoer.org,age30

//函数作为对象，可以作为参数传递，建议将复杂签名定义为函数类型，
//func test(fn func() int) int {
//	return fn()
//}
//// 定义函数类型。
//type FormatFunc func(s string, x, y int) string  //定义函数类型
//
//func format(fn FormatFunc, s string, x, y int) string {  //把复杂的函数类型作为参数，传到函数中
//	return fn(s, x, y)
//}
//func main() {
//	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。
//
//	s2 := format(func(s string, x, y int) string {  //将复制的签名定义函数类型
//		return fmt.Sprintf(s, x, y)
//	}, "%d, %d", 10, 20)
//
//	println(s1, s2)
//}
//输出结果：100 10, 20
//有返回值的函数，必须有明确的终止语句，否则会引发编译错误

//##################函数参数
//函数定义时有参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量。
//
//但当调用函数，传递过来的变量就是函数的实参，函数可以通过两种方式来传递参数：
//值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数

/* 定义相互交换值的函数 */
//func swap(x, y *int) {
//	var temp int
//	temp = *x /* 保存 x 的值 */
//	*x = *y   /* 将 y 值赋给 x */
//	*y = temp /* 将 temp 值赋给 y*/
//
//}
//func main() {
//	var a, b int = 1, 2
//	/*
//	   调用 swap() 函数
//	   &a 指向 a 指针，a 变量的地址
//	   &b 指向 b 指针，b 变量的地址
//	*/
//	swap(&a, &b)
//
//	fmt.Println(a, b)
//}

//注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
//注意2：map、slice、chan、指针、interface默认以引用的方式传递。
//不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
//Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。
//在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上“…”即可。

//func myfunc(args ...int) {    //0个或多个参数
//}
//func add(a int, args…int) int {    //1个或多个参数
//}
//func add(a int, b int, args…int) int {    //2个或多个参数
//}

//注意：其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.
//任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。
//用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的

//func test(s string, n ...int) string {
//	var x int
//	for _, i := range n {
//		x += i
//	}
//
//	return fmt.Sprintf(s, x)
//}
//
//func main() {
//	println(test("sum: %d", 1, 2, 3))
//}


//使用slice做对象变参数时
//func test(s string, n ...int) string {
//	var x int
//	for _, i := range n {
//		x += i
//	}
//
//	return fmt.Sprintf(s, x)
//}
//
//func main() {
//	s := []int{1, 2, 3}
//	res := test("sum: %d", s...)    // slice... 展开slice
//	println(res)
//}


//############函数返回值

//"_"标识符，用来忽略函数的某个返回值，返回值可以呗命名，并且像函数体开头那样被申明使用
//返回值的名称应当具有一定的意义，可以作为文档使用，没有参数的return语句返回各个变量的当前值，这种用法被叫"裸"返回
//直接返回语句仅应当下面短函数中

//func add(a, b int) (c int) {
//	c = a + b
//	return
//}
//func calc(a, b int) (sum int, avg int) {
//	sum = a + b
//	avg = (a + b) / 2
//
//	return
//}
//func main() {
//	var a, b int = 1, 2
//	c := add(a, b)   //调用函数后给c重新赋值
//	sum, avg := calc(a, b)  //通过函数，返回两个变量的值
//	fmt.Println(a, b, c, sum, avg)
//}
//输出结果：1 2 3 3 1

//返回值不能用容器对象接受多返回值，只能用多个变量，或"_"忽略
//func test() (int, int) {
//	return 1, 2
//}
//
//func main() {
//	// s := make([]int, 2)
//	// s = test()   // Error: multiple-value test() in single-value context
//
//	x, _ := test() //忽略一个变量
//	println(x)
//}
//输出结果：1

//多返回值可以其他函数参数调用实惨

//func test() (int, int) {
//	return 1, 2
//}
//func add(x, y int) int {
//	return x + y
//}
//func sum(n ...int) int {
//	var x int
//	for _, i := range n { //i循环等于1，2和为3
//		x += i
//	}
//	return x
//}
//func main() {
//	println(add(test()))
//	println(sum(test()))
//}
//输出结果：3，3


//命名返回参数，可以看作与形参类似的局部变量

//func add(x, y int) (z int) {
//	z = x + y
//	return
//}
//
//func main() {
//	println(add(1, 2))
//}
//输出结果：3


//命名返回参数可被同名局部变量遮蔽，此时需要显式返回。


//func add(x, y int) (z int) {
//	defer func() {  //命名返回参数允许 defer 延迟调用通过闭包读取和修改。
//		z += 100
//	}()
//	z = x + y
//	return
//}
//
//func main() {
//	println(add(1, 2))
//}
//输出结果： 103


//func add(x, y int) (z int) {
//	defer func() {
//		println(z) // 输出: 203
//	}()
//
//	z = x + y
//	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
//}
//
//func main() {
//	println(add(1, 2)) // 输出: 203
//}
//这里注意，执行顺序很重要


//###################匿名函数################################
//匿名函数就是不需要定义函数名的一种函数实现方式，，匿名函数又一个不带函数名的函数申明和函数体组成，匿名函数优越性在于可以直接使用函数内的变量，不必申明

//func main() {
//	getSqrt := func(a float64) float64 {
//		return math.Sqrt(a)
//	}
//	fmt.Println(getSqrt(4))
//}
//输出结果：2


//匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送

//func main() {
//	// --- function variable ---
//	fn := func() { println("Hello, World!") }
//	fn()
//
//	// --- function collection ---
//	fns := [](func(x int) int){
//		func(x int) int { return x + 1 },
//		func(x int) int { return x + 2 },
//	}
//	println(fns[0](100))
//	// --- function as field ---
//	d := struct {
//		fn func() string
//	}{
//		fn: func() string { return "Hello, World!" },
//	}
//	println(d.fn())
//	// --- channel of function ---
//	fc := make(chan func() string, 2)
//	fc <- func() string { return "Hello, World!" }
//	println((<-fc)())
//}

//输出结果
//Hello, World!
//101
//Hello, World!
//Hello, World!

//################闭包和递归###################
//所谓“闭包”，指的是一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。

//func a() func() int {  //闭包
//	i := 0
//	b := func() int {
//		i++
//		fmt.Println(i)
//		return i
//	}
//	return b
//}
//
//func main() {
//	c := a()
//	c()
//	c()
//	c()
//
//	a() //不会输出i
//}
//输出结果1 2 3

//闭包复制的是原对象指针，这就很容易解释延迟引用现象。

//func test() func() {
//	x := 100
//	fmt.Printf("x (%p) = %d\n", &x, x)
//
//	return func() {
//		fmt.Printf("x (%p) = %d\n", &x, x)
//	}
//}
//
//func main() {
//	f := test()
//	f()
//}
//输出结果
//x (0xc0000b2008) = 100
//x (0xc0000b2008) = 100


//test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。当调 匿名函数时，只需以某个寄存器传递该对象即可。

// 外部引用函数参数局部变量
//func add(base int) func(int) int {
//	return func(i int) int {
//		base += i
//		return base
//	}
//}
//
//func main() {
//	tmp1 := add(10)
//	fmt.Println(tmp1(1), tmp1(2))  //返回两个闭包
//	// 此时tmp1和tmp2不是一个实体了
//	tmp2 := add(100)
//	fmt.Println(tmp2(1), tmp2(2))
//}
//输出结果
//11 13
//101 103



// 返回2个函数类型的返回值
//func test01(base int) (func(int) int, func(int) int) {
//	// 定义2个函数，并返回
//	// 相加
//	add := func(i int) int {
//		base += i
//		return base
//	}
//	// 相减
//	sub := func(i int) int {
//		base -= i
//		return base
//	}
//	// 返回
//	return add, sub
//}
//
//func main() {
//	f1, f2 := test01(10)
//	// base一直是没有消
//	fmt.Println(f1(1), f2(2))
//	// 此时base是9
//	fmt.Println(f1(3), f2(4))
//}
//输出结果
//11 9
//12 8



//递归，就是在运行的过程中调用自己。 一个函数调用自己，就叫做递归函数
//1.子问题须与原始问题为同样的事，且更为简单。
//2.不能无限制地调用本身，须有个出口，化简为非递归状况处理。

//一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!。1808年


//func factorial(i int) int {
//	if i <= 1 {
//		return 1
//	}
//	return i * factorial(i-1)
//}
//
//func main() {
//	var i int = 7
//	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
//}
//输出结果
//Factorial of 7 is 5040



//斐波那契数列(Fibonacci)

//这个数列从第3项开始，每一项都等于前两项之和。

//func fibonaci(i int) int {
//	if i == 0 {
//		return 0
//	}
//	if i == 1 {
//		return 1
//	}
//	return fibonaci(i-1) + fibonaci(i-2)
//}
//
//func main() {
//	var i int
//	for i = 0; i < 10; i++ {
//		fmt.Printf("%d\n", fibonaci(i))
//	}
//}
//输出结果
//0
//1
//1
//2
//3
//5
//8
//13
//21
//34
