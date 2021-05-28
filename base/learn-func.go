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


//##############测试函数示例

//单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较。



// split package with a single split function.

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
//func Split(s, sep string) (result []string) {
//	i := strings.Index(s, sep)
//
//	for i > -1 {
//		result = append(result, s[:i])
//		s = s[i+1:]
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}

//在当前目录下，我们创建一个split_test.go的测试文件，并定义一个测试函数如下：


/// split/split_test.go
//
//package split
//
//import (
//"reflect"
//"testing"
//)
//
//func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Split("a:b:c", ":")         // 程序输出的结果
//	want := []string{"a", "b", "c"}    // 期望的结果
//	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//}

//此时包中u如下
//split $ ls -l
//total 16
//-rw-r--r--  1 pprof staff  408  4 29 15:50 split.go
//-rw-r--r--  1 pprof  staff  466  4 29 16:04 split_test.go

//在split包路径下，执行go test命令，可以看到输出结果如下：

//split $ go test
//PASS
//ok      github.com/pprof/studygo/code_demo/test_demo/split       0.005s

//一个测试用例有点单薄，我们再编写一个测试使用多个字符切割字符串的例子，在split_test.go中添加如下测试函数：

//func TestMoreSplit(t *testing.T) {
//	got := Split("abcd", "bc")
//	want := []string{"a", "d"}
//	if !reflect.DeepEqual(want, got) {
//		t.Errorf("excepted:%v, got:%v", want, got)
//	}
//}

//运行go test输出过

//split $ go test
//--- FAIL: TestMultiSplit (0.00s)
//split_test.go:20: excepted:[a d], got:[a cd]
//FAIL
//exit status 1
//FAIL    github.com/pprof/studygo/code_demo/test_demo/split       0.006s

//以为go test命令添加-v参数，查看测试函数名称和运行时间：

//split $ go test -v
//=== RUN   TestSplit
//--- PASS: TestSplit (0.00s)
//=== RUN   TestMoreSplit
//--- FAIL: TestMoreSplit (0.00s)
//split_test.go:21: excepted:[a d], got:[a cd]
//FAIL
//exit status 1
//FAIL    github.com/pprof/studygo/code_demo/test_demo/split       0.005s


//还可以在go test命令后添加-run参数，它对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。

//split $ go test -v -run="More"
//    === RUN   TestMoreSplit
//    --- FAIL: TestMoreSplit (0.00s)
//        split_test.go:21: excepted:[a d], got:[a cd]
//    FAIL
//    exit status 1
//    FAIL    github.com/pprof/studygo/code_demo/test_demo/split       0.006s

//很显然我们最初的split函数并没有考虑到sep为多个字符的情况，我们来修复下这个Bug：

//package split
//
//import "strings"
//
//// split package with a single split function.
//
//// Split slices s into all substrings separated by sep and
//// returns a slice of the substrings between those separators.
//func Split(s, sep string) (result []string) {
//	i := strings.Index(s, sep)
//
//	for i > -1 {
//		result = append(result, s[:i])
//		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}


//注意，当我们修改了我们的代码之后不要仅仅执行那些失败的测试函数，我们应该完整的运行所有的测试，保证不会因为修改代码而引入了新的问题。

//split $ go test -v
//    === RUN   TestSplit
//    --- PASS: TestSplit (0.00s)
//    === RUN   TestMoreSplit
//    --- PASS: TestMoreSplit (0.00s)
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s


//#####################测试组######################

//还想要测试一下split函数对中文字符串的支持，这个时候我们可以再编写一个TestChineseSplit测试函数，但是我们也可以使用如下更友好的一种方式来添加更多的测试用例。

//func TestSplit(t *testing.T) {
//   // 定义一个测试用例类型
//    type test struct {
//        input string
//        sep   string
//        want  []string
//    }
//    // 定义一个存储测试用例的切片
//    tests := []test{
//        {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//        {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//        {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//        {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//    }
//    // 遍历切片，逐一执行测试用例
//    for _, tc := range tests {
//        got := Split(tc.input, tc.sep)
//        if !reflect.DeepEqual(got, tc.want) {
//            t.Errorf("excepted:%v, got:%v", tc.want, got)
//        }
//    }
//}

//代码把多个测试用例合到一起，再次执行go test命令。

//split $ go test -v
//    === RUN   TestSplit
//    --- FAIL: TestSplit (0.00s)
//        split_test.go:42: excepted:[枯藤 树昏鸦], got:[ 枯藤 树昏鸦]
//    FAIL
//    exit status 1
//    FAIL    github.com/pprof/studygo/code_demo/test_demo/split       0.006s


//打印的测试失败提示信息：excepted:[枯藤 树昏鸦], got:[ 枯藤 树昏鸦]，你会发现[ 枯藤 树昏鸦]中有个不明显的空串，这种情况下十分推荐使用%#v的格式化方式。
//
//我们修改下测试用例的格式化输出错误提示部分

//func TestSplit(t *testing.T) {
//   ...
//
//    for _, tc := range tests {
//        got := Split(tc.input, tc.sep)
//        if !reflect.DeepEqual(got, tc.want) {
//            t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//        }
//    }
//}

//运行go test命令后就能看到比较明显的提示信息了：

//    split $ go test -v
//    === RUN   TestSplit
//    --- FAIL: TestSplit (0.00s)
//        split_test.go:42: excepted:[]string{"枯藤", "树昏鸦"}, got:[]string{"", "枯藤", "树昏鸦"}
//    FAIL
//    exit status 1
//    FAIL    github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s

//##############子测试###################


//看起来都挺不错的，但是如果测试用例比较多的时候，我们是没办法一眼看出来具体是哪个测试用例失败了


//func TestSplit(t *testing.T) {
//    type test struct { // 定义test结构体
//        input string
//        sep   string
//        want  []string
//    }
//    tests := map[string]test{ // 测试用例使用map存储
//        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//    }
//    for name, tc := range tests {
//        got := Split(tc.input, tc.sep)
//        if !reflect.DeepEqual(got, tc.want) {
//            t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
//        }
//    }
//}

//同时Go1.7+中新增了子测试，我们可以按照如下方式使用t.Run执行子测试：


//func TestSplit(t *testing.T) {
//    type test struct { // 定义test结构体
//        input string
//        sep   string
//        want  []string
//    }
//    tests := map[string]test{ // 测试用例使用map存储
//        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
//    }
//    for name, tc := range tests {
//        t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
//            got := Split(tc.input, tc.sep)
//            if !reflect.DeepEqual(got, tc.want) {
//                t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//            }
//        })
//    }
//}

//再执行go test命令就能够看到更清晰的输出内容了：

// split $ go test -v
//    === RUN   TestSplit
//    === RUN   TestSplit/leading_sep
//    === RUN   TestSplit/simple
//    === RUN   TestSplit/wrong_sep
//    === RUN   TestSplit/more_sep
//    --- FAIL: TestSplit (0.00s)
//        --- FAIL: TestSplit/leading_sep (0.00s)
//            split_test.go:83: excepted:[]string{"枯藤", "树昏鸦"}, got:[]string{"", "枯藤", "树昏鸦"}
//        --- PASS: TestSplit/simple (0.00s)
//        --- PASS: TestSplit/wrong_sep (0.00s)
//        --- PASS: TestSplit/more_sep (0.00s)
//    FAIL
//    exit status 1
//    FAIL    github.com/pprof/studygo/code_demo/test_demo/split       0.006s

//这个时候我们要把测试用例中的错误修改回来：


//func TestSplit(t *testing.T) {
//    ...
//    tests := map[string]test{ // 测试用例使用map存储
//        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"", "枯藤", "树昏鸦"}},
//    }
//    ...
//}

//通过-run=RegExp来指定运行的测试用例，还可以通过/来指定要运行的子测试用例，例如：go test -v -run=Split/simple只会运行simple对应的子测试用例。


//######################测试覆盖率#################


//测试覆盖率是你的代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。

//Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。例如：

// split $ go test -cover
//    PASS
//    coverage: 100.0% of statements
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       0.005s


//从上面的结果可以看到我们的测试用例覆盖了100%的代码。


//Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。例如：

//    split $ go test -cover -coverprofile=c.out
//    PASS
//    coverage: 100.0% of statements
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       0.005s

//上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。


//######################基准测试###############
//基准测试函数格式


//基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：

//func BenchmarkName(b *testing.B){
//    // ...
//}


//基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。 testing.B拥有的方法如下：

//func (c *B) Error(args ...interface{})
//func (c *B) Errorf(format string, args ...interface{})
//func (c *B) Fail()
//func (c *B) FailNow()
//func (c *B) Failed() bool
//func (c *B) Fatal(args ...interface{})
//func (c *B) Fatalf(format string, args ...interface{})
//func (c *B) Log(args ...interface{})
//func (c *B) Logf(format string, args ...interface{})
//func (c *B) Name() string
//func (b *B) ReportAllocs()
//func (b *B) ResetTimer()
//func (b *B) Run(name string, f func(b *B)) bool
//func (b *B) RunParallel(body func(*PB))
//func (b *B) SetBytes(n int64)
//func (b *B) SetParallelism(p int)
//func (c *B) Skip(args ...interface{})
//func (c *B) SkipNow()
//func (c *B) Skipf(format string, args ...interface{})
//func (c *B) Skipped() bool
//func (b *B) StartTimer()
//func (b *B) StopTimer()


//#################基准测试示例################
//split包中的Split函数编写基准测试如下
//func BenchmarkSplit(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Split("枯藤老树昏鸦", "老")
//	}
//}

//基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试，输出结果如下：

//    split $ go test -bench=Split
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/split
//    BenchmarkSplit-8        10000000               203 ns/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       2.255s


//其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

//我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。

//    split $ go test -bench=Split -benchmem
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/split
//    BenchmarkSplit-8        10000000               215 ns/op             112 B/op          3 allocs/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       2.394s

//其中，112 B/op表示每次操作内存分配了112字节，3 allocs/op则表示每次操作进行了3次内存分配。 我们将我们的Split函数优化如下

//func Split(s, sep string) (result []string) {
//    result = make([]string, 0, strings.Count(s, sep)+1)
//    i := strings.Index(s, sep)
//    for i > -1 {
//        result = append(result, s[:i])
//        s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
//        i = strings.Index(s, sep)
//    }
//    result = append(result, s)
//    return
//}

//这一次我们提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。我们来看一下这个改进会带来多大的性能提升：

//    split $ go test -bench=Split -benchmem
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/split
//    BenchmarkSplit-8        10000000               127 ns/op              48 B/op          1 allocs/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/split       1.423s

//这个使用make函数提前分配内存的改动，减少了2/3的内存分配次数，并且减少了一半的内存分配。

//#############################性能比较函数#############################

//性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下：

//func benchmark(b *testing.B, size int){/* ... */}
//func Benchmark10(b *testing.B){ benchmark(b, 10) }
//func Benchmark100(b *testing.B){ benchmark(b, 100) }
//func Benchmark1000(b *testing.B){ benchmark(b, 1000) }

//编写了一个计算斐波那契数列的函数如下：

//// fib.go
//
//// Fib 是一个计算第n个斐波那契数的函数
//func Fib(n int) int {
//    if n < 2 {
//        return n
//    }
//    return Fib(n-1) + Fib(n-2)
//}


//编写的性能比较函数如下：

//// fib_test.go
//
//func benchmarkFib(b *testing.B, n int) {
//    for i := 0; i < b.N; i++ {
//        Fib(n)
//    }
//}
//
//func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
//func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
//func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
//func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
//func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
//func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//运行基准测试：

//    split $ go test -bench=.
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//    BenchmarkFib1-8         1000000000               2.03 ns/op
//    BenchmarkFib2-8         300000000                5.39 ns/op
//    BenchmarkFib3-8         200000000                9.71 ns/op
//    BenchmarkFib10-8         5000000               325 ns/op
//    BenchmarkFib20-8           30000             42460 ns/op
//    BenchmarkFib40-8               2         638524980 ns/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/fib 12.944s


//这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。

//最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。例如：

//    split $ go test -bench=Fib40 -benchtime=20s
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/pprof/studygo/code_demo/test_demo/fib
//    BenchmarkFib40-8              50         663205114 ns/op
//    PASS
//    ok      github.com/pprof/studygo/code_demo/test_demo/fib 33.849s


//这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了。


//使用性能比较函数做测试的时候一个容易犯的错误就是把b.N作为输入的大小，例如以下两个例子都是错误的示范：

//// 错误示范1
//func BenchmarkFibWrong(b *testing.B) {
//    for n := 0; n < b.N; n++ {
//        Fib(n)
//    }
//}
//
//// 错误示范2
//func BenchmarkFibWrong2(b *testing.B) {
//    Fib(b.N)
//}

//##############################重置时间##################

//b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。

//func BenchmarkSplit(b *testing.B) {
//    time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
//    b.ResetTimer()              // 重置计时器
//    for i := 0; i < b.N; i++ {
//        Split("枯藤老树昏鸦", "老")
//    }
//}

//############并行测试################

//func (b B) RunParallel(body func(PB))会以并行的方式执行给定的基准测试。

//RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。RunParallel通常会与-cpu标志一同使用。

//func BenchmarkSplitParallel(b *testing.B) {
//    // b.SetParallelism(1) // 设置使用的CPU数
//    b.RunParallel(func(pb *testing.PB) {
//        for pb.Next() {
//            Split("枯藤老树昏鸦", "老")
//        }
//    })
//}

//执行一下基准测试：

//split $ go test -bench=.
//goos: darwin
//goarch: amd64
//pkg: github.com/pprof/studygo/code_demo/test_demo/split
//BenchmarkSplit-8                10000000               131 ns/op
//BenchmarkSplitParallel-8        50000000                36.1 ns/op
//PASS
//ok      github.com/pprof/studygo/code_demo/test_demo/split       3.308s


//还可以通过在测试命令后添加-cpu参数如go test -bench=. -cpu 1来指定使用的CPU数量。

//Setup与TearDown

//测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）

//###############TestMain

//通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。

//如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。
//
//一个使用TestMain来设置Setup和TearDown的示例如下：

//func TestMain(m *testing.M) {
//    fmt.Println("write setup code here...") // 测试之前的做一些设置
//    // 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//    retCode := m.Run()                         // 执行测试
//    fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
//    os.Exit(retCode)                           // 退出测试
//}


//需要注意的是：在调用TestMain时, flag.Parse并没有被调用。所以如果TestMain 依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。

//#####################子测试的Setup与Teardown############################

//有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

//// 测试集的Setup与Teardown
//func setupTestCase(t *testing.T) func(t *testing.T) {
//    t.Log("如有需要在此执行:测试之前的setup")
//    return func(t *testing.T) {
//        t.Log("如有需要在此执行:测试之后的teardown")
//    }
//}
//
//// 子测试的Setup与Teardown
//func setupSubTest(t *testing.T) func(t *testing.T) {
//    t.Log("如有需要在此执行:子测试之前的setup")
//    return func(t *testing.T) {
//        t.Log("如有需要在此执行:子测试之后的teardown")
//    }
//}

//使用方式如下：

//func TestSplit(t *testing.T) {
//    type test struct { // 定义test结构体
//        input string
//        sep   string
//        want  []string
//    }
//    tests := map[string]test{ // 测试用例使用map存储
//        "simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
//        "wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
//        "more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
//        "leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"", "枯藤", "树昏鸦"}},
//    }
//    teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
//    defer teardownTestCase(t)            // 测试之后执行testdoen操作
//
//    for name, tc := range tests {
//        t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
//            teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
//            defer teardownSubTest(t)           // 测试之后执行testdoen操作
//            got := Split(tc.input, tc.sep)
//            if !reflect.DeepEqual(got, tc.want) {
//                t.Errorf("excepted:%#v, got:%#v", tc.want, got)
//            }
//        })
//    }
//}


//测试结果如下：

//    split $ go test -v
//    === RUN   TestSplit
//    === RUN   TestSplit/simple
//    === RUN   TestSplit/wrong_sep
//    === RUN   TestSplit/more_sep
//    === RUN   TestSplit/leading_sep
//    --- PASS: TestSplit (0.00s)
//        split_test.go:71: 如有需要在此执行:测试之前的setup
//        --- PASS: TestSplit/simple (0.00s)
//            split_test.go:79: 如有需要在此执行:子测试之前的setup
//            split_test.go:81: 如有需要在此执行:子测试之后的teardown
//        --- PASS: TestSplit/wrong_sep (0.00s)
//            split_test.go:79: 如有需要在此执行:子测试之前的setup
//            split_test.go:81: 如有需要在此执行:子测试之后的teardown
//        --- PASS: TestSplit/more_sep (0.00s)
//            split_test.go:79: 如有需要在此执行:子测试之前的setup
//            split_test.go:81: 如有需要在此执行:子测试之后的teardown
//        --- PASS: TestSplit/leading_sep (0.00s)
//            split_test.go:79: 如有需要在此执行:子测试之前的setup
//            split_test.go:81: 如有需要在此执行:子测试之后的teardown
//        split_test.go:73: 如有需要在此执行:测试之后的teardown
//    === RUN   ExampleSplit
//    --- PASS: ExampleSplit (0.00s)
//    PASS
//    ok      github.com/Q1mi/studygo/code_demo/test_demo/split       0.006s

//#########################示例函数###############################

//###############示例函数的格式

//被go test特殊对待的第三种函数就是示例函数，它们的函数名以Example为前缀。它们既没有参数也没有返回值。标准格式如下：

//func ExampleName() {
//    // ...
//}

//示例函数示例

//下面的代码是我们为Split函数编写的一个示例函数：

//func ExampleSplit() {
//    fmt.Println(split.Split("a:b:c", ":"))
//    fmt.Println(split.Split("枯藤老树昏鸦", "老"))
//    // Output:
//    // [a b c]
//    // [ 枯藤 树昏鸦]
//}


//为你的代码编写示例代码有如下三个用处：

//   示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
//
//    示例函数只要包含了// Output:也是可以通过go test运行的可执行测试。
//
//        split $ go test -run Example
//        PASS
//        ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s
//    示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。下图为strings.ToUpper函数在Playground的示例函数效果。


//###################压力测试########################

//##################Go怎么写测试用例##########

//单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让线上的程序能够在高并发的情况下还能保持稳定
//Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例，
//另外建议安装gotests插件自动生成测试代码:

//go get -u -v github.com/cweill/gotests/...

//#########如何编写测试用例

//由于go test命令只能在一个相应的目录下执行所有文件，所以我们接下来新建一个项目目录gotest,这样我们所有的代码和测试代码都在这个目录下。
//
//接下来我们在该目录下面创建两个文件：gotest.go和gotest_test.go
//
//gotest.go:这个文件里面我们是创建了一个包，里面有一个函数实现了除法运算:


//package gotest
//
//    import (
//        "errors"
//    )
//
//    func Division(a, b float64) (float64, error) {
//        if b == 0 {
//            return 0, errors.New("除数不能为0")
//        }
//
//        return a / b, nil
//    }

//gotest_test.go:这是我们的单元测试文件，但是记住下面的这些原则：
//
//文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
//
//你必须import testing这个包
//
//所有的测试用例函数必须是Test开头
//
//测试用例会按照源代码中写的顺序依次执行
//
//测试函数TestXxx()的参数是testing.T，我们可以使用该类型来记录错误或者是测试状态
//
//测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
//
//函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。


//package gotest
//
//import (
//"testing"
//)

//func Test_Division_1(t *testing.T) {
//	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
//		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
//	} else {
//		t.Log("第一个测试通过了") //记录一些你期望记录的信息
//	}
//}
//
//func Test_Division_2(t *testing.T) {
//	t.Error("就是不通过")
//}
//执行结果

//--- FAIL: Test_Division_2 (0.00 seconds)
//gotest_test.go:16: 就是不通过
//FAIL
//exit status 1
//FAIL    gotest    0.013s

//从这个结果显示测试没有通过，因为在第二个测试函数中我们写死了测试不通过的代码t.Error，那么我们的第一个函数执行的情况怎么样呢？默认情况下执行go test是不会显示测试通过的信息的，我们需要带上参数go test -v，

//    === RUN Test_Division_1
//    --- PASS: Test_Division_1 (0.00 seconds)
//        gotest_test.go:11: 第一个测试通过了
//    === RUN Test_Division_2
//    --- FAIL: Test_Division_2 (0.00 seconds)
//        gotest_test.go:16: 就是不通过
//    FAIL
//    exit status 1
//    FAIL    gotest    0.012s

//上面的输出详细的展示了这个测试的过程，我们看到测试函数1Test_Division_1测试通过，而测试函数2Test_Division_2测试失败了，最后得出结论测试不通过。

//func Test_Division_2(t *testing.T) {
//        if _, e := Division(6, 0); e == nil { //try a unit test on function
//            t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
//        } else {
//            t.Log("one test passed.", e) //记录一些你期望记录的信息
//        }
//    }

//然后我们执行go test -v，就显示如下信息

//    === RUN Test_Division_1
//    --- PASS: Test_Division_1 (0.00 seconds)
//        gotest_test.go:11: 第一个测试通过了
//    === RUN Test_Division_2
//    --- PASS: Test_Division_2 (0.00 seconds)
//        gotest_test.go:20: one test passed. 除数不能为0
//    PASS
//    ok      gotest    0.013s


//########################如何编写压力测试

//压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似,此处不再赘述，但需要注意以下几点：
//
//压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母

//go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表示测试全部的压力测试函数
//
//在压力测试用例中,请记得在循环体内使用testing.B.N,以使测试可以正常的运行 文件名也必须以_test.go结尾
//
//下面我们新建一个压力测试文件webbench_test.go，

//package gotest
//
//import (
//    "testing"
//)
//
//func Benchmark_Division(b *testing.B) {
//    for i := 0; i < b.N; i++ { //use b.N for looping
//        Division(4, 5)
//    }
//}
//
//func Benchmark_TimeConsumingFunction(b *testing.B) {
//    b.StopTimer() //调用该函数停止压力测试的时间计数
//
//    //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
//    //这样这些时间不影响我们测试函数本身的性能
//
//    b.StartTimer() //重新开始时间
//    for i := 0; i < b.N; i++ {
//        Division(4, 5)
//    }
//}
































