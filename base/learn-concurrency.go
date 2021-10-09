package main

import (
	"fmt"
	"runtime"
	"time"
)

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

func a()  {
	for i := 1; i < 10; i ++ {
		fmt.Println("A",i)
	}
}

func b()  {
	for i := 1; i < 10; i++ {
		fmt.Println("B",i)
	}
}

func main()  {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

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










































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































