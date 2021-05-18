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
