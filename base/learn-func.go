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


//###############延迟调用defer
//特性：
//1. 关键字 defer 用于注册延迟调用。
//2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
//3. 多个defer语句，按先进后出的方式执行。
//4. defer语句中的变量，在defer声明时就决定了。

//用途
//1. 关闭文件句柄
//2. 锁资源释放
//3. 数据库连接释放

//defer 是先进后出

//func main() {
//	var whatever [5]struct{}
//
//	for i := range whatever {
//		defer fmt.Println(i)
//	}
//}
//输出结果：
//4
//3
//2
//1
//0

//defer碰上闭包

//func main() {
//	var whatever [5]struct{}
//	for i := range whatever {
//		defer func() { fmt.Println(i) }()
//	}
//}
//输出结果：
//4
//4
//4
//4
//4

//函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.

//type Test struct {
//	name string
//}
//
//func (t *Test) Close() {
//	fmt.Println(t.name, " closed")
//}
//func main() {
//	ts := []Test{{"a"}, {"b"}, {"c"}}
//	for _, t := range ts {
//		defer t.Close()
//	}
//}
//输出结果
//c  closed
//c  closed
//c  closed


//type Test struct {
//	name string
//}
//
//func (t *Test) Close() {
//	fmt.Println(t.name, " closed")
//}
//func Close(t Test) {
//	t.Close()
//}
//func main() {
//	ts := []Test{{"a"}, {"b"}, {"c"}}
//	for _, t := range ts {
//		defer Close(t)
//	}
//}
//输出结果
//c  closed
//b  closed
//a  closed

//type Test struct {
//	name string
//}
//
//func (t *Test) Close() {
//	fmt.Println(t.name, " closed")
//}
//func main() {
//	ts := []Test{{"a"}, {"b"}, {"c"}}
//	for _, t := range ts {
//		t2 := t
//		defer t2.Close()
//	}
//}

//输出结果
//c  closed
//b  closed
//a  closed
//defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说struct这里的this指针如何处理，通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待



//多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。

//func test(x int) {
//	defer println("a")
//	defer println("b")
//
//	defer func() {
//		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
//	}()
//
//	defer println("c")
//}
//
//func main() {
//	test(0)
//}
//输出结果
//c
//b
//a
//panic: runtime error: integer divide by zero

//*延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。


//func test() {
//	x, y := 10, 20
//
//	defer func(i int) {
//		println("defer:", i, y) // y 闭包引用
//	}(x) // x 被复制
//
//	x += 10
//	y += 100
//	println("x =", x, "y =", y)
//}
//
//func main() {
//	test()
//}

//输出结果
//x = 20 y = 120
//defer: 10 120


//*滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
//var lock sync.Mutex
//
//func test() {
//	lock.Lock()
//	lock.Unlock()
//}
//
//func testdefer() {
//	lock.Lock()
//	defer lock.Unlock()
//}
//
//func main() {
//	func() {
//		t1 := time.Now()
//
//		for i := 0; i < 10000; i++ {
//			test()
//		}
//		elapsed := time.Since(t1)
//		fmt.Println("test elapsed: ", elapsed)
//	}()
//	func() {
//		t1 := time.Now()
//
//		for i := 0; i < 10000; i++ {
//			testdefer()
//		}
//		elapsed := time.Since(t1)
//		fmt.Println("testdefer elapsed: ", elapsed)
//	}()
//
//}
//输出结果

//test elapsed:  142.132µs
//testdefer elapsed:  147.136µ



//############defer陷阱

//func foo(a, b int) (i int, err error) {
//	defer fmt.Printf("first defer err %v\n", err)
//	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
//	defer func() { fmt.Printf("third defer err %v\n", err) }()
//	if b == 0 {
//		err = errors.New("divided by zero!")
//		return
//	}
//
//	i = a / b
//	return
//}
//
//func main() {
//	foo(2, 0)
//}
//输出结果
//third defer err divided by zero!
//second defer err <nil>
//first defer err <nil>
//解释：如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。


//func foo() (i int) {
//
//	i = 0
//	defer func() {
//		fmt.Println(i)
//	}()
//
//	return 2
//}
//
//func main() {
//	foo()
//}
//输出结果：2
//解释：在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。所以defer closure 输出结果为 2 而不是 1。


//###########defer nil函数
//func test() {
//	var run func() = nil
//	defer run()
//	fmt.Println("runs")
//}
//
//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//	test()
//}

//输出结果
//runs
//runtime error: invalid memory address or nil pointer dereference

//解释：名为 test 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。然而值得注意的是，run() 的声明是没有问题，因为在test函数运行完成后它才会被调用。


//在错误的位置使用defer

//当http.Get失败时抛出异常
//func do() error {
//	res, err := http.Get("http://www.google.com")
//	defer res.Body.Close()
//	if err != nil {
//		return err
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func main() {
//	do()
//}

//输出结果
//panic: runtime error: invalid memory address or nil pointer dereference
//[signal SIGSEGV: segmentation violation code=0x1 addr=0x40 pc=0x1247f2b]

//因为在这里我们并没有检查我们的请求是否成功执行，当它失败的时候，我们访问了 Body 中的空变量 res ，因此会抛出异常


//当且仅当 http.Get 成功执行时才使用 defer
//func do() error {
//	res, err := http.Get("http://www.wulaoer.org")
//	if res != nil {
//		defer res.Body.Close()
//	}
//
//	if err != nil {
//		return err
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func main() {
//	do()
//}

//当有错误的时候，err 会被返回，否则当整个函数返回的时候，会关闭 res.Body 。

//解释：在这里，你同样需要检查 res 的值是否为 nil ，这是 http.Get 中的一个警告。通常情况下，出错的时候，返回的内容应为空并且错误会被返回，可当你获得的是一个重定向 error 时， res 的值并不会为 nil ，但其又会将错误返回。上面的代码保证了无论如何 Body 都会被关闭，如果你没有打算使用其中的数据，那么你还需要丢弃已经接收的数据。


//############不检查错误
//f.Close() 可能会返回一个错误，可这个错误会被我们忽略掉

//func do() error {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//
//	if f != nil {
//		defer f.Close()
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func main() {
//	do()
//}


//改进一下

//func do() error {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//
//	if f != nil {
//		defer func() {
//			if err := f.Close(); err != nil {
//				// log etc
//			}
//		}()
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func main() {
//	do()
//}


//通过命名的返回变量来返回 defer 内的错误。


//func do() (err error) {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//
//	if f != nil {
//		defer func() {
//			if ferr := f.Close(); ferr != nil {
//				err = ferr
//			}
//		}()
//	}
//
//	// ..code...
//
//	return nil
//}
//
//func main() {
//	do()
//}


//释放相同的资源


//如果你尝试使用相同的变量释放不同的资源，那么这个操作可能无法正常执行。

//func do() error {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func() {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close book.txt err %v\n", err)
//			}
//		}()
//	}
//
//	// ..code...
//
//	f, err = os.Open("another-book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func() {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close another-book.txt err %v\n", err)
//			}
//		}()
//	}
//
//	return nil
//}
//
//func main() {
//	do()
//}
//输出结果： defer close book.txt err close ./another-book.txt: file already closed
//当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (another-book.txt)。而且两个 defer 都会将这个资源作为最后的资源来关闭


//func do() error {
//	f, err := os.Open("book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func(f io.Closer) {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close book.txt err %v\n", err)
//			}
//		}(f)
//	}
//
//	// ..code...
//
//	f, err = os.Open("another-book.txt")
//	if err != nil {
//		return err
//	}
//	if f != nil {
//		defer func(f io.Closer) {
//			if err := f.Close(); err != nil {
//				fmt.Printf("defer close another-book.txt err %v\n", err)
//			}
//		}(f)
//	}
//
//	return nil
//}
//
//func main() {
//	do()
//}


//###################异常处理
//golang没有结构化异常，使用panic抛出错误，recover捕获错误
//Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

//panic

//1、内置函数
//2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
//3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
//4、直到goroutine整个退出，并报告错误


//recover


//1、内置函数
//2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
//3、一般的调用建议
//a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
//b). 可以获取通过panic传递的error

//注意

//1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
//2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
//3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。


//func main() {
//	test()
//}
//
//func test() {
//	defer func() {
//		if err := recover(); err != nil {
//			println(err.(string)) // 将 interface{} 转型为具体类型。
//		}
//	}()
//
//	panic("panic error!")
//}

//输出结果
//panic error!


//由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
//    func panic(v interface{})
//    func recover() interface{}


//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println(err)
//		}
//	}()
//
//	var ch chan int = make(chan int, 10)
//	close(ch)
//	ch <- 1
//}
//输出结果
//send on closed channel

//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。

//func test() {
//	defer func() {
//		fmt.Println(recover())
//	}()
//
//	defer func() {
//		panic("defer panic")
//	}()
//
//	panic("test panic")
//}
//
//func main() {
//	test()
//}
//输出结果
//defer panic

//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。

//func test() {
//	defer func() {
//		fmt.Println(recover()) //有效
//	}()
//	defer recover()              //无效！
//	defer fmt.Println(recover()) //无效！
//	defer func() {
//		func() {
//			println("defer inner")
//			recover() //无效！
//		}()
//	}()
//
//	panic("test panic")
//}
//
//func main() {
//	test()
//}

//输出结果
//defer inner
//<nil>
//test panic


//使用延迟匿名函数或下面这样都是有效的。


//func except() {
//	fmt.Println(recover())
//}
//
//func test() {
//	defer except()
//	panic("test panic")
//}
//
//func main() {
//	test()
//}
//输出结果
//test panic

//如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执 。

//func test(x, y int) {
//	var z int
//
//	func() {
//		defer func() {
//			if recover() != nil {
//				z = 0
//			}
//		}()
//		panic("test panic")
//		z = x / y
//		return
//	}()
//
//	fmt.Printf("x / y = %d\n", z)
//}
//
//func main() {
//	test(2, 1)
//}

//输出结果
//x / y = 0


//除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。


//标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。

//var ErrDivByZero = errors.New("division by zero")
//
//func div(x, y int) (int, error) {
//	if y == 0 {
//		return 0, ErrDivByZero
//	}
//	return x / y, nil
//}
//
//func main() {
//	defer func() {
//		fmt.Println(recover())
//	}()
//	switch z, err := div(10, 0); err {
//	case nil:
//		println(z)
//	case ErrDivByZero:
//		panic(err)
//	}
//}
//输出结果
//division by zero


//Go实现类似 try catch 的异常处理

//func Try(fun func(), handler func(interface{})) {
//	defer func() {
//		if err := recover(); err != nil {
//			handler(err)
//		}
//	}()
//	fun()
//}
//
//func main() {
//	Try(func() {
//		panic("test panic")
//	}, func(err interface{}) {
//		fmt.Println(err)
//	})
//}

//输出结果
//test panic


//惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error



//############################单元测试#####################################

//go test工具

//go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中

//在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
//
//类型	格式	作用
//测试函数	函数名前缀为Test	测试程序的一些逻辑行为是否正确
//基准函数	函数名前缀为Benchmark	测试函数的性能
//示例函数	函数名前缀为Example	为文档提供示例文档

//go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件

//Golang单元测试对文件名和方法名，参数都有很严格的要求。
//1、文件名必须以xx_test.go命名
//2、方法必须是Test[^a-z]开头
//3、方法参数必须 t *testing.T
//4、使用go test执行单元测试

//go test是go语言自带的测试工具，其中包含的是两类，单元测试和性能测试
//格式形如： go test [-c] [-i] [build flags] [packages] [flags for test binary]
//参数解读：
//-c : 编译go test成为可执行的二进制文件，但是不运行测试。
//-i : 安装测试包依赖的package，但是不运行测试。
//关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空
//关于packages，调用go help packages，这些是关于包的管理，一般设置为空
//关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数
//-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。
//-test.run pattern: 只跑哪些单元测试用例
//-test.bench patten: 只跑那些性能测试用例
//-test.benchmem : 是否在性能测试的时候输出内存情况
//-test.benchtime t : 性能测试运行的时间，默认是1s
//-test.cpuprofile cpu.out : 是否输出cpu性能分析文件
//-test.memprofile mem.out : 是否输出内存性能分析文件
//-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件
//-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。
//你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。
//-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下
//-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。
//-test.timeout t : 如果测试用例运行时间超过t，则抛出panic
//-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理
//-test.short : 将那些运行时间较长的测试用例运行时间缩短


//目录结构：
//test
//|
//—— calc.go
//|
//—— calc_test.go


//######测试函数

//测试函数格式

//每个测试函数必须导入testing包，测试函数的基本格式（签名）如下

//func TestName(t *testing.T){
//	// ...
//}

//测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子：\

//func TestAdd(t *testing.T){ ... }
//func TestSum(t *testing.T){ ... }
//func TestLog(t *testing.T){ ... }

//其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：

//func (c *T) Error(args ...interface{})
//func (c *T) Errorf(format string, args ...interface{})
//func (c *T) Fail()
//func (c *T) FailNow()
//func (c *T) Failed() bool
//func (c *T) Fatal(args ...interface{})
//func (c *T) Fatalf(format string, args ...interface{})
//func (c *T) Log(args ...interface{})
//func (c *T) Logf(format string, args ...interface{})
//func (c *T) Name() string
//func (t *T) Parallel()
//func (t *T) Run(name string, f func(t *T)) bool
//func (c *T) Skip(args ...interface{})
//func (c *T) SkipNow()
//func (c *T) Skipf(format string, args ...interface{})
//func (c *T) Skipped() bool



























