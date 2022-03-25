//package main
//
//import (
//	"fmt"
//	"time"
//)

//###################并发编程####################

//进程和线程

//A. 进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
//    B. 线程是进程的一个执行实体,是CPU调度和分派的基本单位,它是比进程更小的能独立运行的基本单位。
//    C.一个进程可以创建和撤销多个线程;同一个进程中的多个线程之间可以并发执行。

//并发和并行

//   A. 多线程程序在一个核的cpu上运行，就是并发。
//    B. 多线程程序在多个核的cpu上运行，就是并行。

//协程和线程

//协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
//线程：一个线程上可以跑多个协程，协程是轻量级的线程。

//goroutine 只是由官方实现的超级"线程池"

//每个实力4~5KB的栈内存占用和由于实现机制而大幅减少的创建和销毁开销是go高并发的根本原因

//并发不是并行

//并发主要由切换时间片来实现"同时"运行，并行则是直接利用多核实现多线程的运行，go可以设置使用核数，以发挥多核计算机的能力

//#####################1. Goroutine#############################

//Go语言中的goroutine就是这样一种机制，goroutine的概念类似于线程，但 goroutine是由Go的运行时（runtime）调度和管理的。Go程序会智能地将 goroutine 中的任务合理地分配给每个CPU。Go语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。
//
//在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个goroutine去执行这个函数就可以了，就是这么简单粗暴。

//启动单个goroutine
//启动goroutine的方式非常简单，只需要在调用的函数（普通函数和匿名函数）前面加上一个go关键字。

//func hello() {
//	fmt.Println("Hello Goroutine!")
//}
//func main() {
//	hello()
//	fmt.Println("main goroutine done!")
//}
//
//输出结果：
//Hello Goroutine!
//main goroutine done!


//func hello() {
//	fmt.Println("Hello Goroutine!")
//}
//func main() {
//	go hello()
//	fmt.Println("main goroutine done!")
//}
//
//输出结果：
//main goroutine done!
//Hello Goroutine!

//在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。

//当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束，main函数所在的goroutine就像是权利的游戏中的夜王，其他的goroutine都是异鬼，夜王一死它转化的那些异鬼也就全部GG了。

//如果想让main函数等一下hello函数，最简单的方法就是time.Sleep

//func hello() {
//	fmt.Println("Hello Goroutine!")
//}
//func main() {
//	go hello()
//	fmt.Println("main goroutine done!")
//	time.Sleep(time.Second)
//}
//输出结果：
//main goroutine done!
//Hello Goroutine!

//这是因为启动一个gorountine的时候需要时间，而main函数继续执行,main函数结束，gorountine也会结束

//启动多个goroutine

//在go中并发就是启动多个goroutne，看下面的例子

//var wg sync.WaitGroup
//
//func hello(i int)  {
//	defer wg.Done()
//	fmt.Println("hello Goroutine!!!",i)
//}
//func main()  {
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go hello(i)
//	}
//	wg.Wait()
//}

//数据结果
//hello Goroutine!!! 9
//hello Goroutine!!! 0
//hello Goroutine!!! 7
//hello Goroutine!!! 8
//hello Goroutine!!! 4
//hello Goroutine!!! 2
//hello Goroutine!!! 1
//hello Goroutine!!! 5
//hello Goroutine!!! 3
//hello Goroutine!!! 6

//这是因为10个goroutine是并发执行的，而goroutine的调度是随机的

//注意：如果主协程退出了，其他任务也会结束，看下面的例子：

//func main()  {
//	//合起来
//	go func() {
//		i:=0
//		for {
//			i++
//			fmt.Printf("new goroutine: i = %d\n",i)
//			time.Sleep(time.Second)
//		}
//	}()
//	i:=0
//	//主协程
//	for {
//		i++
//		fmt.Printf("main goroutine: i = %d\n",i)
//		time.Sleep(time.Second)
//		if i == 2 {
//			break
//		}
//	}
//}

//输出结果
//main goroutine: i = 1
//new goroutine: i = 1
//new goroutine: i = 2
//main goroutine: i = 2

//主协程终止，任务也相应的结束

//#############gorountine与线程

//###可增长的栈
//OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这个大。所以在Go语言中一次创建十万左右的goroutine也是可以的。

//####goroutine调度

//GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。

//1.G很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息

//2.P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P的队列里抢任务。

//3.M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；

//P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。

//P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失

//单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。

//###runtime包

//runtime.Gosched()用于让出CPU时间片，这就像跑接力赛，A跑了一会碰到代码runtime.Gosched()就把接力棒交给B了，A歇着了，B继续跑。

//func main()  {
//	go func(s string) {
//		for i := 0; i < 100; i++ {
//			fmt.Println(s)
//		}
//	}("word")
//	for i := 0;i <2; i++ {
//		//runtime.GOMAXPROCS(4)
//		runtime.Gosched()
//
//		fmt.Println("hello")
//	}
//}

//输出结果

//word
//word
//hello
//hello

//这里需要注意，输出结果中，先输出了word，后输出的hello，，这是因为主线程退出了，上面的函数就没有机会在执行了，这里主要是未来输出结果简短才这样做的，如果把word的for循环改成100，你会发现word的循环没执行完就结束了
//如果代码中通过runtime.GOMAXPROCS(n) 其中n是整数，指定使用多核，goroutins都是在一个线程里，他们之间通过不停的让出时间片轮流运行，

//runtime.Goexit() 退出当前协程

//func main()  {
//	go func() {
//		defer fmt.Println("this is defer A")
//		func() {
//			defer fmt.Println("ehis is defer B")
//			runtime.Goexit()
//			fmt.Println("b")
//
//		}()
//		fmt.Println("A")
//	}() //别忘记()
//	//防止阻塞
//	for{}
//}

//输出结果
//ehis is defer B
//this is defer A

//调用runtime.goExit()将立即终止当前goroutine执行，调度器确保所有已注册defer延迟调度被执行。


//return结束当前函数,并返回指定值
//runtime.Goexit结束当前goroutine,其他的goroutine不受影响,主程序也一样继续运行
//os.Exit会结束当前程序,不管你三七二十一

//####runtime.GOMAXPROCS

//Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数，例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上（GOMAXPROCS是m:n调度中的n）

//Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。

//Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。

//可以通过将任务分配到不同到CPU逻辑核心上实现并行到效果

//func a()  {
//	for i := 1; i < 10; i ++ {
//		fmt.Println("A",i)
//	}
//}
//
//func b()  {
//	for i := 1; i < 10; i++ {
//		fmt.Println("B",i)
//	}
//}
//
//func main()  {
//	runtime.GOMAXPROCS(1)
//	go a()
//	go b()
//	time.Sleep(time.Second)
//}
//输出结果：
//B 1
//B 2
//B 3
//B 4
//B 5
//B 6
//B 7
//B 8
//B 9
//A 1
//A 2
//A 3
//A 4
//A 5
//A 6
//A 7
//A 8
//A 9

//这是把两个核心任务分配到一个CPU上了，做完一个任务再做另一个任务，下面将逻辑核心数设为2

//func a()  {
//	for i := 1; i < 10; i ++ {
//		fmt.Println("A",i)
//	}
//}
//
//func b()  {
//	for i := 1; i < 10; i++ {
//		fmt.Println("B",i)
//	}
//}
//
//func main()  {
//	runtime.GOMAXPROCS(2)
//	go a()
//	go b()
//	time.Sleep(time.Second)
//}

//输出结果
//B 1
//B 2
//B 3
//B 4
//B 5
//A 1
//A 2
//A 3
//A 4
//A 5
//A 6
//A 7
//A 8
//A 9
//B 6
//B 7
//B 8
//B 9

//Go语言中的操作系统线程和goroutine的关系
//1、一个操作系统线程对应用户态多个goroutine
//2、go程序可以同时使用多个操作系统线程
//3、goroutine和OS线程是多对多的关系，即m：n

// ### Channel

//channel，其实单纯的对将函数并发执行是没有任何意义对，函数与函数之间需要交换数据才能提醒并发函数对意义，虽然可以使用共享内存进行数据交换，但是共享内存在不同对goroutine中容易发生竟态问题，为了保证数据交换对正确性，必须使用互斥量对内存进行加锁，这种做法势必造成性能问题

//Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信

//如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制

//Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

//channel是一种类型，一种引用类型。声明通道类型的格式如下：

//var 变量 chan 元素类型

//举例
//var ch1 chan int   // 声明一个传递整型的通道
//var ch2 chan bool  // 声明一个传递布尔型的通道
//var ch3 chan []int // 声明一个传递int切片的通道

//通道是引用类型，通道类型的空值是nil。

//func main()  {
//	var ch chan int
//	fmt.Println(ch)
//}
//输出结果<nil>

//声明的通道后需要使用make函数初始化之后才能使用

//创建channel的格式如下：

//make(chan 元素类型, [缓冲大小])  channel的缓冲大小是可选的。

//举例

//ch4 := make(chan int)
//ch5 := make(chan bool)
//ch6 := make(chan []int)

//channel操作

//通道有发送（send）、接收(receive）和关闭（close）三种操作

//发送和接收都使用<-符号。

//现在我们先使用以下语句定义一个通道：

//ch := make(chan int)

// 发送
//将一个值发送到通道中
//ch <- 10 // 把10发送到ch中

//接收
//从一个通道中接收值
//x := <- ch // 从ch中接收值并赋值给变量x
//<-ch       // 从ch中接收值，忽略结果

//关闭
//我们通过调用内置的close函数来关闭通道。
// close(ch)
//关于关闭通道需要注意的事情是，只有在通知接收方goroutine所有的数据都发送完毕的时候才需要关闭通道。通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的

//关闭后的通道有以下特点：
//1.对一个关闭的通道再发送值就会导致panic。
//2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//4.关闭一个已经关闭的通道会导致panic。

//无缓冲的通道又称为阻塞的通道
//func main()  {
//	ch := make(chan int)
//	ch <- 10
//	fmt.Println("发送成功")
//}

//因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，简单来说就是无缓冲的通道必须有接收才能发送
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan send]:
//main.main()
//	/Users/wolf/Documents/GitHub/golang-learn/base/learn-concurrency.go:391 +0x54

//上面的代码会阻塞在ch <- 10这一行代码形成死锁

//func recv(c chan int)  {
//	ret := <- c //接受
//	fmt.Println("接受成功",ret)
//}
//
//func main()  {
//	ch := make(chan  int) //初始化
//	go recv(ch) // 启用goroutine从通道接收值
//	ch <- 10 //发送
//	fmt.Println("发送成功")
//}

//输出结果
//接受成功 10
//发送成功

//无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

//使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。

//有缓冲的通道

//make函数初始化通道的时候为其指定通道的容量，例如：

//func main()  {
//	ch := make(chan int,1)
//	ch <- 10
//	fmt.Println("发送成功")
//}

//只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。

//内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

//  close()

//可以通过内置的close()函数关闭channel（如果你的管道不往里存值或者取值的时候一定记得关闭管道）

//func main()  {
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//		close(c)
//	}()
//	for {
//		if data, ok := <-c; ok {
//			fmt.Println(data)
//		} else {
//			break
//		}
//	}
//	fmt.Println("main结束")
//}
//输出结果
//0
//1
//2
//3
//4
//main结束

//如何优雅的从通道循环取值

//当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。那如何判断一个通道是否被关闭了呢？

//我们来看下面这个例子：

//chan 练习

//func main()  {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go func() {
//		for i := 0; i < 10; i++ {
//			ch1 <- i
//		}
//		close(ch1)
//	}()
//	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	go func() {
//		for {
//			i,ok := <- ch1 //// 通道关闭后再取值ok=false
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//		}
//		close(ch2)
//	}()
//	// 在主goroutine中从ch2中接收值打印
//	for i := range ch2 { // 通道关闭后会退出for range循环
//		fmt.Println(i)
//	}
//}

//输出结果
//0
//1
//4
//9
//16
//25
//36
//49
//64
//81

//从上面的例子中我们看到有两种方式在接收值的时候判断通道是否被关闭，我们通常使用的是for range的方式。

//  单向通道

//有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收

//Go语言中提供了单向通道来处理这种情况。例如，我们把上面的例子改造如下：

//func counter(out chan<- int)  {
//	for i := 0; i < 10; i++ {
//		out <- i
//	}
//	close(out)
//}
//func squarer(out chan<- int, in <- chan  int)  {
//	for i := range in {
//		out <- i * i
//	}
//	close(out)
//}
//func printer(in <-chan int)  {
//	for i := range in {
//		fmt.Println(i)
//	}
//}
//
//func main()  {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	go counter(ch1)
//	go squarer(ch2,ch1)
//	printer(ch2)
//}

//其中
//    1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
//    2.<-chan int是一个只能接收的通道，可以接收但是不能发送。


//输出结果
//0
//1
//4
//9
//16
//25
//36
//49
//64
//81

//在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来是不可以的

//通道总结

//channel常见的异常总结，如下图：

//                           chanel异常情况总结
// channel   nil   非空   空的      满了          没满
// 接收     阻塞    接收值 阻塞      接收值         接收值
// 发送    阻塞    发送值  发送值    阻塞           发送值
// 关闭   panic   关闭成功 关闭成功， 关闭成功后，   关闭成功，
//               读完数据  返回零值  读完数据后     读完数据后
//               后返回零值         返回零值       返回零值

//注意:关闭已经关闭的channel也会引发panic。

//   Goroutine池

// worker pool（goroutine池）
//本质上是生产者消费者模型
//可以有效控制goroutine数量，防止暴涨
//需求：
//计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//随机生成数字进行计算
//控制台输出结果如下：
//job id:2088376 randnum:3117708862702414273 result:73
//job id:2088377 randnum:8121940470451593634 result:76
//job id:2088378 randnum:6311301581954619591 result:78
//job id:2088379 randnum:375513122517260052 result:57
//job id:2088380 randnum:8753223451094272103 result:68
//job id:2088381 randnum:8607478679035457116 result:94
//job id:2088382 randnum:5583163151852913213 result:72
//job id:2088383 randnum:6101000195477271840 result:63
//job id:2088384 randnum:3471443027424726698 result:83

//import (
//	"fmt"
//	"math/rand"
//)
//
//type Job struct {
//	// id
//	Id int
//	// 需要计算的随机数
//	RandNum int
//}
//
//type Result struct {
//	// 这里必须传对象实例
//	Job *Job
//	// 求和
//	sum int
//}
//// 创建工作池
//// 参数1：开几个协程
//func createPool(num int, JobChan chan *Job, resultChan chan *Result)  {
//	for i := 0; i < num; i++ {
//		go func(jonChan chan *Job, resultChan chan *Result) {
//			// 执行运算
//			// 遍历job管道所有数据，进行相加
//			for job := range  jonChan {
//				// 随机数接过来
//				r_num := job.RandNum
//				// 随机数每一位相加
//				// 定义返回值
//				var sum int
//				for r_num != 0 {
//					tmp := r_num % 10
//					sum += tmp
//					r_num /= 10
//				}
//				// 想要的结果是Result
//				r := &Result{
//					job: job,
//					sum: sum,
//				}
//				//运算结果扔到管道
//				resultChan <- r
//			}
//		}(JobChan,resultChan)
//	}
//}
//
//
//func main()  {
//	// 需要2个管道
//	// 1.job管道
//	jobChan := make(chan *Job, 128)
//	// 2.结果管道
//	resultChan := make(chan *Result, 128)
//	// 3.创建工作池
//	createPool(64, jobChan,resultChan)
//	// 4.开个打印的协程
//	go func(resultChan chan *Result) {
//		// 遍历结果管道打印
//		for result := range resultChan {
//			fmt.Printf("job id:%v randnum:%v result:%d\n",result.job.Id,result.job.RandNum,result.sum)
//		}
//	}(resultChan)
//	var id int
//	// 循环创建job，输入到管道
//	for {
//		id++
//		// 生成随机数
//		r_num := rand.Int()
//		job := &Job{
//			Id:      id,
//			RandNum: r_num,
//		}
//		jobChan <- job
//	}
//}


//输出结果

//job id:2088376 randnum:3117708862702414273 result:73
//job id:2088377 randnum:8121940470451593634 result:76
//job id:2088378 randnum:6311301581954619591 result:78
//job id:2088379 randnum:375513122517260052 result:57
//job id:2088380 randnum:8753223451094272103 result:68
//job id:2088381 randnum:8607478679035457116 result:94
//job id:2088382 randnum:5583163151852913213 result:72
//job id:2088383 randnum:6101000195477271840 result:63
//job id:2088384 randnum:3471443027424726698 result:83
//job id:2088385 randnum:2994512928832232024 result:77

//这是一个循环，所以输出结果就截一部分把


// 定时器

//Time 时间到了，执行只执行一次


//func main() {
//	// 1.timer基本使用
//	//timer1 := time.NewTimer(2 * time.Second)
//	//t1 := time.Now()
//	//fmt.Printf("t1:%v\n", t1)
//	//t2 := <-timer1.C
//	//fmt.Printf("t2:%v\n", t2)
//
//	// 2.验证timer只能响应1次
//	//timer2 := time.NewTimer(time.Second)
//	//for {
//	// <-timer2.C
//	// fmt.Println("时间到")
//	//}
//
//	// 3.timer实现延时的功能
//	//(1)
//	//time.Sleep(time.Second)
//	//(2)
//	//timer3 := time.NewTimer(2 * time.Second)
//	//<-timer3.C
//	//fmt.Println("2秒到")
//	//(3)
//	//<-time.After(2*time.Second)
//	//fmt.Println("2秒到")
//
//	// 4.停止定时器
//	//timer4 := time.NewTimer(2 * time.Second)
//	//go func() {
//	// <-timer4.C
//	// fmt.Println("定时器执行了")
//	//}()
//	//b := timer4.Stop()
//	//if b {
//	// fmt.Println("timer4已经关闭")
//	//}
//
//	// 5.重置定时器
//	timer5 := time.NewTimer(3 * time.Second)
//	timer5.Reset(1 * time.Second)
//	fmt.Println(time.Now())
//	fmt.Println(<-timer5.C)
//	for {
//	}
//}
//输出结果
//2021-10-11 10:18:04.712797 +0800 CST m=+0.000111222
//2021-10-11 10:18:05.716333 +0800 CST m=+1.003695397

//func main() {
//	// 1.获取ticker对象
//	ticker := time.NewTicker(1 * time.Second)
//	i := 0
//	// 子协程
//	go func() {
//		for {
//			//<-ticker.C
//			i++
//			fmt.Println(<-ticker.C)
//			if i == 5 {
//				//停止
//				ticker.Stop()
//			}
//		}
//	}()
//	for {   //为了让定时任务持续运行，到预设到时间执行上面到语句
//
//	}
//}

//输出结果
//2021-10-11 10:21:48.342021 +0800 CST m=+1.004599546
//2021-10-11 10:21:49.339176 +0800 CST m=+2.001802235
//2021-10-11 10:21:50.337893 +0800 CST m=+3.000565954
//2021-10-11 10:21:51.341476 +0800 CST m=+4.004224053
//2021-10-11 10:21:52.338478 +0800 CST m=+5.001245642

//   select

//select多路复用，在某中场景中，需要同时多个通道接收数据，同在在接收数据时，如果没有数据可以接收将会发送阻塞，看下面的遍历

//for{
//    // 尝试从ch1接收值
//    data, ok := <-ch1
//    // 尝试从ch2接收值
//    data, ok := <-ch2
//    …
//}

//这种方式虽然可以接收多个通道需求，但是性能较差，为了应对这种环境，go内置了select关键字，可以同时响应多个通道的操作

//select的使用类似switch语句，可以有一系列的case分支和默认的分支，每个case会对应一个通道的通信(接收或发送)过程，select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句。具体如下：

//    select {
//    case <-chan1:
//       // 如果chan1成功读到数据，则进行该case处理语句
//    case chan2 <- 1:
//       // 如果成功向chan2写入数据，则进行该case处理语句
//    default:
//       // 如果上面都没有成功，则进入default处理流程
//    }

//select可以同时监听一个或多个channel，直到其中一个channel ready

package main

//func test1(ch chan string)  {
//	time.Sleep(time.Second * 5)
//	ch <- "test1"
//}
//func test2(ch chan string)  {
//	time.Sleep(time.Second * 2)
//	ch <- "test2"
//}
//
//func main()  {
//	// 2个管道
//	output1 := make(chan string)
//	outout2 := make(chan string)
//	//跑2个子协程，写数据
//	go test1(output1)
//	go test2(outout2)
//	// 用select监控
//	select {
//	case s1 := <-output1:
//		fmt.Println("s1=",s1)
//	case s2 := <-outout2:
//		fmt.Println("s2",s2)
//
//
//	}
//}

//输出结果
//s2 test2

// 如果有多个channel同时ready，则随机选择一个执行

//func main()  {
//	int_chan := make(chan int,1)
//	string_chan := make(chan string,1)
//	go func() {
//		int_chan <- 1
//	}()
//	go func() {
//		string_chan <- "hello"
//	}()
//	select {
//	case value := <- int_chan:
//		fmt.Println("int",value)
//	case value := <- string_chan:
//		fmt.Println("string",value)
//	}
//	fmt.Println("main结束")
//}

//输出结果

//string hello
//main结束

//可以用于判断管道是否存满

//判断管道是否存满
//func main()  {
//	//创建管道
//	output1 := make(chan string,10)
//	// 子协程写数据
//	go write(output1)
//	//取数据
//	for s := range output1 {
//		fmt.Println("res",s)
//		time.Sleep(time.Second)
//	}
//}
//
//func write(ch chan string)  {
//	for {
//		// 写数据
//		select {
//		case ch <- "hello":
//			fmt.Println("write hello")
//		default:
//			fmt.Println("channel full")
//		}
//		time.Sleep(time.Millisecond * 500)
//	}
//}

//  并发锁和安全锁

//有时多个goroutine同时操作一个资源，这种情况就会发生竟态问题(数据竟态).例如：现实生活中的例子有十字路口被各个方向的的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

//var x int64
//var wg sync.WaitGroup
//
//func add()  {
//	for i := 0; i < 10; i++ {
//		x = x + 1
//		fmt.Println("这是x的值",x)
//	}
//	wg.Done()
//}
//
//func main()  {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

//输出结果

//这是x的值 1
//这是x的值 3
//这是x的值 4
//这是x的值 5
//这是x的值 6
//这是x的值 7
//这是x的值 8
//这是x的值 9
//这是x的值 10
//这是x的值 11
//这是x的值 2
//这是x的值 12
//这是x的值 13
//这是x的值 14
//这是x的值 15
//这是x的值 16
//这是x的值 17
//这是x的值 18
//这是x的值 19
//这是x的值 20
//20

//上面的代码开启了两个goroutine去累加遍历x的值，这两个goroutine在访问和修改x遍历的时候就会存在数据竞争，导致最后的结果与期待的不符。

//  互斥锁

//互斥锁是一种常用的控制共享资源访问的方法，她能够保证同时只有一个goroutine可以访问共享资源，go语言中使用sync包的Mutex类型来实现互斥锁，使用互斥锁来修复上面代码的问题：


//var x int64
//var wg sync.WaitGroup
//var lock sync.Mutex
//
//func add()  {
//	for i := 0; i <10; i++ {
//		lock.Lock()  //加锁
//		x = x + 1
//		fmt.Println("这个是x的值",x)
//		lock.Unlock() //解锁
//	}
//	wg.Done()
//}
//
//func main()  {
//	wg.Add(2)
//	go add()
//	go add()
//	wg.Wait()
//	fmt.Println(x)
//}

//输出结果

//这个是x的值 1
//这个是x的值 2
//这个是x的值 3
//这个是x的值 4
//这个是x的值 5
//这个是x的值 6
//这个是x的值 7
//这个是x的值 8
//这个是x的值 9
//这个是x的值 10
//这个是x的值 11
//这个是x的值 12
//这个是x的值 13
//这个是x的值 14
//这个是x的值 15
//这个是x的值 16
//这个是x的值 17
//这个是x的值 18
//这个是x的值 19
//这个是x的值 20
//20


// 使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁，当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略时随机的。

//   读写互斥锁

//互斥锁是完全的互斥，但是有些环境是读多写少的，当我们并发的去读取资源不涉及资源修改的时候是没有必要加锁的，这种场景中使用读写锁是更好的一种选择，读写锁在Go语言中使用sync包中的RWMutex类型

//读写锁分为两种，读锁和写锁，当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待，当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待

//读写锁实例

//var (
//	x    int64
//	wg   sync.WaitGroup
//	lock sync.Mutex
//	rwlock sync.RWMutex
//)
//
//func write()  {
//	// lock.Lock()   //加互斥锁
//	rwlock.Lock()    //加写锁
//	x = x + 1
//	time.Sleep(10 * time.Millisecond)  //假设读操作耗时10毫秒
//	rwlock.Unlock()     //解写锁
//	wg.Done()     //解互斥锁
//}
//
//func read()  {
//	//local.Lock()   //加互斥锁
//	rwlock.RLock()     //加读锁
//	time.Sleep(time.Millisecond)   //假设读操作耗时1毫秒
//	rwlock.RUnlock()      //解读锁
//	wg.Done()    //解互斥锁
//}
//
//func main()  {
//	start := time.Now()
//	for i := 0; i < 10; i ++ {
//		wg.Add(1)
//		go write()
//	}
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go read()
//	}
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//}

//输出结果
//112.646527ms

//需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁读又是就发挥不出来了。


//   Sync

//sync.WaitGroup 在代码中生硬的使用time.Sleep肯定不合适，go语言中可以使用sync，WaitGroup来实现并发任务的同步，sync.WaitGroup有以下几个方法：

//              方法名                                   功能
// (wg * WaitGroup)Add(delta int)                计数器+delta
// (wg *WaitGroup)Done()                         计数器-1
// (wg *WaitGroup)Wait()                         阻塞直到计算器变为0

//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少，例如当启动N个并发任务时，就将计数器值增加N，每个任务完成时通过调用Done()方法将计数器减1。通过调用Wait()来等待并发任务执行完，当计数器为0时，并使所有任务并发任务已经完成。

//下面看利用sync.WaitGroup将下面当代码优化以下

//var wg sync.WaitGroup
//
//func hello()  {
//	defer wg.Done()
//	fmt.Println("hello groutine!")
//}
//func main()  {
//	wg.Add(1)
//	go hello()
//	fmt.Println("main goroutine done!")
//	wg.Wait()
//}
//输出结果

//main goroutine done!
//hello groutine!

//需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。

//    sync.Once

// 在编程中，我们需要确保某些操作在搞并发场景中只执行一次，例如只加载一次配置文件，只关闭一次通道等。

//go语言中的sync包中提供了一个针对只执行一次场景的解决方案-sync.Once

//sync.Once 只有一个Do方法。其签名如下：

//  func (o *Once) Do(f func()) {}

// 注意，如果执行的函数f需要传递参数，就需要搭配闭包来使用

//  加载配置文件示例

//延迟一个开销很大的初始化操作到真正用到到时候在执行是一个很好到时间，因为预先初始化一个变量(比如在init函数中完成初始化)会正价程序到启动耗时，二期有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做到，我们看以下例子：

//var icons map[string]image.Image
//
//func loadIcons()  {
//	icons = map[string]image.Image{
//		"left": loadIcon("left.png"),
//		"up": loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down": loadIcon("down.png"),
//	}
//}
//
//// Icon 被多个goroutine调用时不是并发安全的
//
//func Icon(name string) image.Image {
//	if icons == nil {
//		loadIcons()
//	}
//	return icons[name]
//}

// 多个goroutine并发调用Icon函数时不是并发安全的，现代的编译器和CPU可能会在保证每个goroutne都满足串行一致的基础上自由地重排访问内存的顺序，IoadIcons函数可能会被重排为以下结果：

//func loadIcons() {
//	icons = make(map[string]image.Image)
//	icons["left"] = loadIcon("left.png")
//	icons["up"] = loadIcon("up.png")
//	icons["right"] = loadIcon("right.png")
//	icons["down"] = loadIcon("down.png")
//}

//在这种情况下就会出现即使判断了icons不是nil也不意味着变量初始化完成了。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化icons的时候不会被其他的goroutine操作，但是这样做又会引发性能问题

//使用sync.Once改造的示例代码如下：

//var  icons map[string]image.Image
//
//func loadIcons() {
//	icons = map[string]image.Image{
//		"left":  loadIcon("left.png"),
//		"up":    loadIcon("up.png"),
//		"right": loadIcon("right.png"),
//		"down":  loadIcon("down.png"),
//	}
//}
//
//// Icon 是并发安全的
//func Icon(name string) image.Image {
//	loadIconsOnce.Do(loadIcons)
//	return icons[name]
//}

//sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。

//    sync.Map

//go语言中内置的map不是并发安全的，看下面的示例

//var m = make(map[string]int)
//
//func get(key string) int {
//	return m[key]
//}
//func set(key string,value int)  {
//	m[key] = value
//}
//
//func main()  {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key,n)
//			fmt.Printf("k=:%v,v:=%v\n",key,get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

//输出结果

//k=:0,v:=0
//k=:19,v:=19
//k=:9,v:=9
//k=:6,v:=6
//k=:13,v:=13
//k=:14,v:=14
//k=:1,v:=1
//k=:2,v:=2
//k=:7,v:=7
//k=:12,v:=12
//fatal error: concurrent map writes

//上面的代码开启少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误

//像这种场景下就需要为map加锁来保证并发的安全性了，Go语言的sync包中提供了一个开箱即用的并发安全版map–sync.Map。开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。

//var m = sync.Map{}
//
//func main()  {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			m.Store(key,n)
//			value, _ := m.Load(key)
//			fmt.Printf("k=:%v,v:=%v\n",key,value)
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

//输出结果

//k=:19,v:=19
//k=:11,v:=11
//k=:0,v:=0
//k=:2,v:=2
//k=:14,v:=14
//k=:3,v:=3
//k=:15,v:=15
//k=:1,v:=1
//k=:4,v:=4
//k=:5,v:=5
//k=:6,v:=6
//k=:10,v:=10
//k=:7,v:=7
//k=:9,v:=9
//k=:16,v:=16
//k=:13,v:=13
//k=:12,v:=12
//k=:17,v:=17
//k=:18,v:=18
//k=:8,v:=8

//    原子操作(atomic包)

// 原子操作 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。Go语言中原子操作由内置的标准库sync/atomic提供

//方法                                                        	解释
//func LoadInt32(addr int32) (val int32)
//func LoadInt64(addr `int64) (val int64)<br>func LoadUint32(addruint32) (val uint32)<br>func LoadUint64(addruint64) (val uint64)<br>func LoadUintptr(addruintptr) (val uintptr)<br>func LoadPointer(addrunsafe.Pointer`) (val unsafe.Pointer)	读取操作
//func StoreInt32(addr *int32, val int32)                     写入操作
//func StoreInt64(addr *int64, val int64)
//func StoreUint32(addr *uint32, val uint32)
//func StoreUint64(addr *uint64, val uint64)
//func StoreUintptr(addr *uintptr, val uintptr)
//func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
//func AddInt32(addr *int32, delta int32) (new int32)            修改操作
//func AddInt64(addr *int64, delta int64) (new int64)
//func AddUint32(addr *uint32, delta uint32) (new uint32)
//func AddUint64(addr *uint64, delta uint64) (new uint64)
//func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
//func SwapInt32(addr *int32, new int32) (old int32)                交换操作
//func SwapInt64(addr *int64, new int64) (old int64)
//func SwapUint32(addr *uint32, new uint32) (old uint32)
//func SwapUint64(addr *uint64, new uint64) (old uint64)
//func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
//func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
//func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)            比较并交换操作
//func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
//func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
//func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
//func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
//func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)



//我们填写一个示例来比较下互斥锁和原子操作的性能。

//var x int64
//var l sync.Mutex
//var wg sync.WaitGroup
//
//// 普通版加函数
//func add() {
//	// x = x + 1
//	x++ // 等价于上面的操作
//	wg.Done()
//}
//
//// 互斥锁版加函数
//func mutexAdd() {
//	l.Lock()
//	x++
//	l.Unlock()
//	wg.Done()
//}
//
//// 原子操作版加函数
//func atomicAdd() {
//	atomic.AddInt64(&x, 1)
//	wg.Done()
//}
//
//func main() {
//	start := time.Now()
//	for i := 0; i < 10000; i++ {
//		wg.Add(1)
//		// go add()       // 普通版add函数 不是并发安全的
//		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
//		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
//	}
//	wg.Wait()
//	end := time.Now()
//	fmt.Println(x)
//	fmt.Println(end.Sub(start))
//}


//输出结果

//10000
//4.567658ms

//     GMP 原理与调度

//一、Golang “调度器” 的由来？

//一切的软件都是跑在操作系统上，真正用来干活 (计算) 的是 CPU。早期的操作系统每个程序就是一个进程，知道一个程序运行完，才能进行下一个进程，就是 “单进程时代”

//早期的单进程操作系统，面临 2 个问题：

//1。单一的执行流程，计算机只能一个任务一个任务处理。
//2。进程阻塞所带来的 CPU 时间浪费。

//那么能不能有多个进程来宏观一起来执行多个任务呢？
//
//后来操作系统就具有了最早的并发能力：多进程并发，当一个进程阻塞的时候，切换到另外等待执行的进程，这样就能尽量把 CPU 利用起来，CPU 就不浪费了。

//在多进程 / 多线程的操作系统中，就解决了阻塞的问题，因为一个进程阻塞 cpu 可以立刻切换到其他进程中去执行，而且调度 cpu 的算法可以保证在运行的进程都可以被分配到 cpu 的运行时间片。这样从宏观来看，似乎多个进程是在同时被运行。

//但新的问题就又出现了，进程拥有太多的资源，进程的创建、切换、销毁，都会占用很长的时间，CPU 虽然利用起来了，但如果进程过多，CPU 有很大的一部分都被用来进行进程调度了。

//怎么才能提高 CPU 的利用率呢？
//
//但是对于 Linux 操作系统来讲，cpu 对进程的态度和线程的态度是一样的。

//很明显，CPU 调度切换的是进程和线程。尽管线程看起来很美好，但实际上多线程开发设计会变得更加复杂，要考虑很多同步竞争等问题，如锁、竞争冲突等。

//(3) 协程来提高 CPU 利用率

//多进程、多线程已经提高了系统的并发能力，但是在当今互联网高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存 (进程虚拟内存会占用 4GB [32 位操作系统], 而线程也要大约 4MB)

//大量的进程 / 线程出现了新的问题
//
//高内存占用
//调度的高消耗 CPU

//好了，然后工程师们就发现，其实一个线程分为 “内核态 “线程和” 用户态 “线程。
//
//一个 “用户态线程” 必须要绑定一个 “内核态线程”，但是 CPU 并不知道有 “用户态线程” 的存在，它只知道它运行的是一个 “内核态线程”(Linux 的 PCB 进程控制块)。

//​ 看到这里，我们就要开脑洞了，既然一个协程 (co-routine) 可以绑定一个线程 (thread)，那么能不能多个协程 (co-routine) 绑定一个或者多个线程 (thread) 上呢。

//N 个协程绑定 1 个线程，优点就是协程在用户态线程即完成切换，不会陷入到内核态，这种切换非常的轻量快速。但也有很大的缺点，1 个进程的所有协程都绑定在 1 个线程上
//
//缺点：
//
//某个程序用不了硬件的多核加速能力
//
//一旦某协程阻塞，造成线程阻塞，本进程的其他协程都无法执行了，根本就没有并发的能力了。

//1:1 关系
//
//1 个协程绑定 1 个线程，这种最容易实现。协程的调度都由 CPU 完成了，不存在 N:1 缺点，
//
//缺点：
//
//协程的创建、删除和切换的代价都由 CPU 完成，有点略显昂贵了。

//M 个协程绑定 1 个线程，是 N:1 和 1:1 类型的结合，克服了以上 2 种模型的缺点，但实现起来最为复杂。

// 协程跟线程是有区别的，线程由 CPU 调度是抢占式的，协程由用户态调度是协作式的，一个协程让出 CPU 后，才执行下一个协程。

//(4) Go 语言的协程 goroutine

//Go 为了提供更容易使用的并发方法，使用了 goroutine 和 channel。goroutine 来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被 runtime 调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。
//
//Go 中，协程被称为 goroutine，它非常轻量，一个 goroutine 只占几 KB，并且这几 KB 就足够 goroutine 运行完，这就能在有限的内存空间内支持大量 goroutine，支持了更多的并发。虽然一个 goroutine 的栈只占几 KB，但实际是可伸缩的，如果需要更多内容，runtime 会自动为 goroutine 分配。
//
//Goroutine 特点：
//
//占用内存更小（几 kb）
//调度更灵活 (runtime 调度)

//(5) 被废弃的 goroutine 调度器
//​好了，既然我们知道了协程和线程的关系，那么最关键的一点就是调度协程的调度器的实现了。
//
//Go 目前使用的调度器是 2012 年重新设计的，因为之前的调度器性能存在问题，所以使用 4 年就被废弃了，那么我们先来分析一下被废弃的调度器是如何运作的？
//
//大部分文章都是会用 G 来表示 Goroutine，用 M 来表示线程，那么我们也会用这种表达的对应关系。

//Processor，它包含了运行 goroutine 的资源，如果线程想运行 goroutine，必须先获取 P，P 中还包含了可运行的 G 队列。

//(1) GMP 模型

//在 Go 中，线程是运行 goroutine 的实体，调度器的功能是把可运行的 goroutine 分配到工作线程上

//全局队列（Global Queue）：存放等待运行的 G。
//P 的本地队列：同全局队列类似，存放的也是等待运行的 G，存的数量有限，不超过 256 个。新建 G’时，G’优先加入到 P 的本地队列，如果队列满了，则会把本地队列中一半的 G 移动到全局队列。
//P 列表：所有的 P 都在程序启动时创建，并保存在数组中，最多有 GOMAXPROCS(可配置) 个。
//M：线程想运行任务就得获取 P，从 P 的本地队列获取 G，P 队列为空时，M 也会尝试从全局队列拿一批 G 放到 P 的本地队列，或从其他 P 的本地队列偷一半放到自己 P 的本地队列。M 运行 G，G 执行之后，M 会从 P 获取下一个 G，不断重复下去。

//Goroutine 调度器和 OS 调度器是通过 M 结合起来的，每个 M 都代表了 1 个内核线程，OS 调度器负责把内核线程分配到 CPU 的核上执行。

//有关 P 和 M 的个数问题
//
//1、P 的数量：
//
//由启动时环境变量 $GOMAXPROCS 或者是由 runtime 的方法 GOMAXPROCS() 决定。这意味着在程序执行的任意时刻都只有 $GOMAXPROCS 个 goroutine 在同时运行。
//2、M 的数量:
//
//go 语言本身的限制：go 程序启动时，会设置 M 的最大数量，默认 10000. 但是内核很难支持这么多的线程数，所以这个限制可以忽略。
//runtime/debug 中的 SetMaxThreads 函数，设置 M 的最大数量
//一个 M 阻塞了，会创建新的 M。

//M 与 P 的数量没有绝对关系，一个 M 阻塞，P 就会去创建或者切换另一个 M，所以，即使 P 的默认数量是 1，也有可能会创建很多个 M 出来。

// P 和 M 何时会被创建

//P 何时创建：在确定了 P 的最大数量 n 后，运行时系统会根据这个数量创建 n 个 P。
//
//2、M 何时创建：没有足够的 M 来关联 P 并运行其中的可运行的 G。比如所有的 M 此时都阻塞住了，而 P 中还有很多就绪任务，就会去寻找空闲的 M，而没有空闲的，就会去创建新的 M。



















































































































































































































































































































