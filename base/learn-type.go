package main

import "fmt"

//golang没有"类"的概念，也不支持"类"的继承等面向对象的概念，golang中通过结构体的内嵌在配合接口比面向对象具有更高的扩展性和灵活性。
//golan中，基本的数据类型，如string，整型，浮点型，布尔等数据类型，可以使用type关键字来定义类型
//将MyInt定义为int类型
//type MyInt int
//除了int的类型，还有byte，rune等类型，类型别名与类型定义表面上看是只有一个等号等差异，但是下面看看：

//类型定义
type NewInt int
//类型别名
type MyInt = int
func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt  //结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型
	fmt.Printf("type of b:%T\n", b) //type of b:int     //b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
}
//输出结果：
//type of a:main.NewInt
//type of b:int

















