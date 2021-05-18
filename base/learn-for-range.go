package main

//range类似迭代器，返回(索引,值)或(键，值)
//for循环的range格式可以对slice，map，数组，字符串等进行迭代循环，格式如下：
//for key, value := range oldMap {
//	newMap[key] = value
//}



//func main() {
//	s := "abc"
//	// 忽略 2nd value，支持 string/array/slice/map。
//	for i := range s {
//		println("aaaa",s[i])
//	}
//	// 忽略 index。
//	for _, c := range s {
//		println("ccccc",c)
//	}
//	// 忽略全部返回值，仅迭代。
//	for range s {
//	}
//	m := map[string]int{"a": 1, "b": 2}
//	// 返回 (key, value)。
//	for k, v := range m {
//		println(k, v)
//	}
//}
//输出结果
//aaaa 97
//aaaa 98
//aaaa 99
//ccccc 97
//ccccc 98
//ccccc 99
//a 1
//b 2

//range复制对象
//func main() {
//	a := [3]int{0, 1, 2}
//	for i, v := range a { // index、value 都是从复制品中取出。
//		if i == 0 { // 在修改前，我们先修改原数组。
//			a[1], a[2] = 999, 999
//			fmt.Println(a) // 确认修改有效，输出 [0, 999, 999]。
//		}
//		a[i] = v + 100 // 使用复制品中取出的 value 修改原数组。
//	}
//	fmt.Println(a) // 输出 [100, 101, 102]。
//}
//输出结果：
//[0 999 999]
//[100 101 102]


//func main() {
//	s := []int{1, 2, 3, 4, 5}
//	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
//		if i == 0 {
//			s = s[:3]  // 对 slice 的修改，不会影响 range。
//			s[2] = 100 // 对底层数据的修改。
//		}
//		println(i, v)
//	}
//}
//输出结果
//0 1
//1 2
//2 100
//3 4
//4 5



func main() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		//if i == 0 {
		//	s = s[:3]  // 对 slice 的修改，不会影响 range。
		//	s[2] = 100 // 对底层数据的修改。
		//}
		println(i, v)
	}
}
//输出结果
//0 1
//1 2
//2 3
//3 4
//4 5

//for 和 for range有什么区别?
//主要是使用场景不同
//for可以
//遍历array和slice
//遍历key为整型递增的map
//遍历string
//for range可以完成所有for可以做的事情，却能做到for不能做的，包括
//遍历key为string类型的map并同时获取key和value
//遍历channel