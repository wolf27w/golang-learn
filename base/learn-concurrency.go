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

























































