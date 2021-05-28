package main

import "fmt"

//############方法定义

//只能为当前包内命名类型定义方法。
//• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
//• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
//• 不支持方法重载，receiver 只是参数签名的组成部分。
//• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。

//一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
//
//所有给定类型的方法属于该类型的方法集。


//########方法定义

//func (recevier type) methodName(参数列表)(返回值列表){}
//
//参数和返回值可以省略

//package main
//
//type Test struct{}
//
//// 无参数、无返回值
//func (t Test) method0() {
//
//}
//
//// 单参数、无返回值
//func (t Test) method1(i int) {
//
//}
//
//// 多参数、无返回值
//func (t Test) method2(x, y int) {
//
//}
//
//// 无参数、单返回值
//func (t Test) method3() (i int) {
//    return
//}
//
//// 多参数、多返回值
//func (t Test) method4(x, y int) (z int, err error) {
//    return
//}
//
//// 无参数、无返回值
//func (t *Test) method5() {
//
//}
//
//// 单参数、无返回值
//func (t *Test) method6(i int) {
//
//}
//
//// 多参数、无返回值
//func (t *Test) method7(x, y int) {
//
//}
//
//// 无参数、单返回值
//func (t *Test) method8() (i int) {
//    return
//}
//
//// 多参数、多返回值
//func (t *Test) method9(x, y int) (z int, err error) {
//    return
//}
//
//func main() {}


//下面定义一个结构体类型和该类型的一个方法：





//结构体
//type User struct {
//	Name  string
//	Email string
//}
//
////方法
//func (u User) Notify() {
//	fmt.Printf("%v : %v \n", u.Name, u.Email)
//}
//func main() {
//	// 值类型调用方法
//	u1 := User{"golang", "golang@golang.com"}
//	u1.Notify()
//	// 指针类型调用方法
//	u2 := User{"go", "go@go.com"}
//	u3 := &u2
//	u3.Notify()
//}

//输出结果

//golang : golang@golang.com
//go : go@go.com

//注释：首先我们定义了一个叫做 User 的结构体类型，然后定义了一个该类型的方法叫做 Notify，该方法的接受者是一个 User 类型的值。要调用 Notify 方法我们需要一个 User 类型的值或者指针。


//Go 调整和解引用指针使得调用可以被执行。注意，当接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。

//我们修改 Notify 方法，让它的接受者使用指针类型：

//结构体
//type User struct {
//	Name  string
//	Email string
//}
//
////方法
//func (u *User) Notify() {
//	fmt.Printf("%v : %v \n", u.Name, u.Email)
//}
//func main() {
//	// 值类型调用方法
//	u1 := User{"golang", "golang@golang.com"}
//	u1.Notify()
//	// 指针类型调用方法
//	u2 := User{"go", "go@go.com"}
//	u3 := &u2
//	u3.Notify()
//}

//输出结果
//golang : golang@golang.com
//go : go@go.com


//注意：当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作。

//方法不过是一种特殊的函数，只需将其还原，就知道 receiver T 和 *T 的差别。

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

func main() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)
}
//输出结果
//Data: 0xc0000a8008
//Value: 0xc0000a8018
//Pointer: 0xc0000a8008
//Value: 0xc0000a8020
//Pointer: 0xc0000a8008











































