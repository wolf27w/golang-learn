package main

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

//这是因为启动一个gorountine的时候需要时间，而main函数继续执行


























































































































































































































































































































































































































































































































