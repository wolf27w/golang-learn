package main

//golang的循环语句
//for循环语法
//for init; condition; post { }
//for condition { }
//for { }
//init： 一般为赋值表达式，给控制变量赋初值；
//condition： 关系表达式或逻辑表达式，循环控制条件；
//post： 一般为赋值表达式，给控制变量增量或减量。
//for语句执行过程如下：
//①先对表达式 init 赋初值；
//②判别赋值表达式 init 是否满足给定 condition 条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
//举例

//func main()  {
//	s := "abc"
//	for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
//		println(s[i])
//	}
//	n := len(s)
//	for n > 0 {                // 替代 while (n > 0) {}
//		println(s[n])        // 替代 for (; n > 0;) {}
//		n--
//	}
//	for {                    // 替代 while (true) {}
//		println(s)            // 替代 for (;;) {}
//	}
//}


//func length(s string) int {
//	println("call length.")
//	return len(s)
//}
//
//func main() {
//	s := "abcd"
//
//	for i, n := 0, length(s); i < n; i++ {     // 避免多次调用 length 函数。
//		println(i, s[i])
//	}
//}
//输出结果：
//call length.
//0 97
//1 98
//2 99
//3 100

//for多个循环
//func main() {
//	var b int = 15
//	var a int
//	numbers := [6]int{1, 2, 3, 5}
//	/* for 循环 */
//	for a := 0; a < 10; a++ {
//		fmt.Printf("a 的值为: %d\n", a)
//	}
//	for a < b {
//		a++
//		fmt.Printf("a 的值为: %d\n", a)
//	}
//	for i,x:= range numbers {
//		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
//	}
//}