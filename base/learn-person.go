package main

//################面向对象

//go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//    go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段

//人
//type Person struct {
//	name string
//	sex  string
//	age  int
//}
//
//type Student struct {
//	Person
//	id   int
//	addr string
//}
//
//func main() {
//	// 初始化
//	s1 := Student{Person{"5lmh", "man", 20}, 1, "bj"}
//	fmt.Println(s1)
//
//	s2 := Student{Person: Person{"5lmh", "man", 20}}
//	fmt.Println(s2)
//
//	s3 := Student{Person: Person{name: "5lmh"}}
//	fmt.Println(s3)
//}
//输出结果

//{{5lmh man 20} 1 bj}
//{{5lmh man 20} 0 }
//{{5lmh  0} 0 }



