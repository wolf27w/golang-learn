package main

import "fmt"

//流程控制语句if，通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
//if语法
//if 布尔表达式 {
///* 在布尔表达式为 true 时执行 */
//}
//if 判读语句

//func main()  {
//	x := 0
//	if n := "abc";x > 0 {
//		fmt.Println(n[2])
//	} else if x < 0 {
//		fmt.Println(n[1])
//	} else {
//		fmt.Println(n[0]) //字符串a的二进制值97
//	}
//}
// 这是一个多判读语句，要不大于0，要不小于0或者等于0

//func main() {
//	/* 定义局部变量 */
//	var a int = 10
//	/* 使用 if 语句判断布尔表达式 */
//	if a < 20 {
//		/* 如果条件为 true 则执行以下语句 */
//		fmt.Printf("a 小于 20\n" )
//	}
//	fmt.Printf("a 的值为 : %d\n", a)  //这里只有两个条件判读，要不大于要不不大于
//}


//if..else 多条件，多个判读

//func main() {
//	/* 局部变量定义 */
//	var a int = 100
//	/* 判断布尔表达式 */
//	if a < 20 {
//		/* 如果条件为 true 则执行以下语句 */
//		fmt.Printf("a 小于 20\n" )
//	} else {
//		/* 如果条件为 false 则执行以下语句 */
//		fmt.Printf("a 不小于 20\n" )
//	}
//	fmt.Printf("a 的值为 : %d\n", a)
//
//}



//if..else嵌套
//if 布尔表达式 1 {
///* 在布尔表达式 1 为 true 时执行 */
//if 布尔表达式 2 {
///* 在布尔表达式 2 为 true 时执行 */
//}
//}


func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	/* 判断条件 */
	if a == 100 {
		/* if 条件语句为 true 执行 */
		if b == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
		}
	}
	fmt.Printf("a 值为 : %d\n", a )
	fmt.Printf("b 值为 : %d\n", b )
}
