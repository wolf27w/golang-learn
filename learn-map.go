//package main
//
//import "fmt"
//
//func main()  {
//	slice := []int{0,1,2,3}
//	m := make(map[int]*int)
//	for key,val := range slice {
//		value := val
//		m[key] = &value
//	}
//	for k,v := range m {
//		fmt.Println(k,"--->",*v)
//	}
//}

package main

//map中的数据都是成对出现的，语法结构为：map[KeyType]ValueType keyType表示键类型，valueType表示键对应的值的类型
//map类型的变量默认初始值为nil，需要使用make()函数累分配内存，语法为： make(map[KeyType]ValueType, [cap]  cap表示map的容量，改参数虽然不是必须的，但是在初始化map的时候就为其制定一个合适的容量
//func main() {
//	scoreMap := make(map[string]int, 8)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	fmt.Println(scoreMap)
//	fmt.Println(scoreMap["小明"])
//	fmt.Printf("type of a:%T\n", scoreMap)
//}
//输出结果：
//map[小明:100 张三:90]
//100
//type of a:map[string]int


///########################
//map支持在生命的时候填充元素

//func main() {
//	userInfo := map[string]string{
//		"username": "pprof.cn",
//		"password": "123456",
//	}
//	fmt.Println(userInfo) //map[password:123456 username:pprof.cn]
//}

//###############################
//判断map中键是否存在特殊的写法，格式：value, ok := map[key]

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
//	v, ok := scoreMap["张三"]
//	if ok {
//		fmt.Println(v) //输出结果90
//	} else {
//		fmt.Println("查无此人")
//	}
//}

//#############################
//使用rang遍历map

//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["王五"] = 60
//	for k, v := range scoreMap {
//		fmt.Println(k, v)
//	}
//}
//输出结果：
//张三 90
//小明 100
//王五 60
//只想遍历key
//func main() {
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["王五"] = 60
//	for k := range scoreMap {
//		fmt.Println(k)
//	}
//}
//输出结果：
//小明
//王五
//张三

//############################delete()函数
//delete()函数删除键值对，语法：delete(map, key) mao表示要删除的map，key表示要删除的键值对的键

//func main(){
//	scoreMap := make(map[string]int)
//	scoreMap["张三"] = 90
//	scoreMap["小明"] = 100
//	scoreMap["王五"] = 60
//	delete(scoreMap, "小明")//将小明:100从map中删除
//	for k,v := range scoreMap{
//		fmt.Println(k, v)
//	}
//}
//输出结果：
//张三 90
//王五 60

//##############################
//根据指定顺序遍历map
//func main() {
//	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
//
//	var scoreMap = make(map[string]int, 200)
//
//	for i := 0; i < 100; i++ {
//		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
//		value := rand.Intn(100)          //生成0~99的随机整数
//		scoreMap[key] = value
//	}
//	//取出map中的所有key存入切片keys
//	var keys = make([]string, 0, 200)
//	for key := range scoreMap {
//		keys = append(keys, key)
//	}
//	//对切片进行排序
//	sort.Strings(keys)
//	//按照排序后的key遍历map
//	for _, key := range keys {
//		fmt.Println(key, scoreMap[key])
//	}
//}
//输出结果：
//stu00 8
//stu01 64
//stu02 82
//stu03 34
//stu04 8
//stu05 15
//stu06 83
//stu07 89
//stu08 1
//stu09 51
//stu10 73
//stu11 92
//stu12 8
//stu13 26
//stu14 87
//stu15 85
//stu16 27
//stu17 90
//stu18 72
//stu19 94
//stu20 64
//stu21 2
//stu22 65
//stu23 19
//stu24 10
//stu25 30
//stu26 62
//stu27 94
//stu28 58
//stu29 71
//stu30 69
//stu31 78
//stu32 2
//stu33 90
//stu34 39
//stu35 93
//stu36 17
//stu37 2
//stu38 65
//stu39 12
//stu40 67
//stu41 81
//stu42 90
//stu43 43
//stu44 0
//stu45 90
//stu46 97
//stu47 44
//stu48 89
//stu49 83
//stu50 52
//stu51 90
//stu52 29
//stu53 92
//stu54 78
//stu55 65
//stu56 33
//stu57 25
//stu58 48
//stu59 18
//stu60 24
//stu61 95
//stu62 2
//stu63 88
//stu64 6
//stu65 37
//stu66 58
//stu67 78
//stu68 31
//stu69 77
//stu70 99
//stu71 80
//stu72 47
//stu73 12
//stu74 5
//stu75 90
//stu76 72
//stu77 75
//stu78 32
//stu79 7
//stu80 35
//stu81 27
//stu82 20
//stu83 39
//stu84 80
//stu85 40
//stu86 1
//stu87 78
//stu88 61
//stu89 1
//stu90 52
//stu91 19
//stu92 33
//stu93 75
//stu94 43
//stu95 79
//stu96 94
//stu97 93
//stu98 36
//stu99 54

//#########元素为map类型的切片
//func main() {
//	var mapSlice = make([]map[string]string, 3)
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//	fmt.Println("after init")
//	// 对切片中的map元素进行初始化
//	mapSlice[0] = make(map[string]string, 10)
//	mapSlice[0]["name"] = "王五"
//	mapSlice[0]["password"] = "123456"
//	mapSlice[0]["address"] = "红旗大街"
//	for index, value := range mapSlice {
//		fmt.Printf("index:%d value:%v\n", index, value)
//	}
//}
//输出结果：
//index:0 value:map[]
//index:1 value:map[]
//index:2 value:map[]
//after init
//index:0 value:map[address:红旗大街 name:王五 password:123456]
//index:1 value:map[]
//index:2 value:map[]

//#####值为切片类型的map

//func main() {
//	var sliceMap = make(map[string][]string, 3)
//	fmt.Println(sliceMap)
//	fmt.Println("after init")
//	key := "中国"
//	value, ok := sliceMap[key]
//	if !ok {
//		value = make([]string, 0, 2)
//	}
//	value = append(value, "北京", "上海")
//	sliceMap[key] = value
//	fmt.Println(sliceMap)
//}
//输出结果：
//map[]
//after init
//map[中国:[北京 上海]]



//func main() {
//
//slice := []int{0,1,2,3}
//m := make(map[int]*int)
//fmt.Println()
//for key,val := range slice {
//	value := val
//	m[key] = &value
//}
//
//for k,v := range m {
//	fmt.Println(k,"===>",*v)
//}
//}




