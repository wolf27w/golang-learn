//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"regexp"
//)
//
////这个只是一个简单的版本只是获取QQ邮箱并且没有进行封装操作，另外爬出来的数据也没有进行去重操作
//var (
//	// \d是数字
//	reQQEmail = `(\d+)@qq.com`
//)
//
//// 爬邮箱
//func GetEmail() {
//	// 1.去网站拿数据
//	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
//	HandleError(err, "http.Get url")
//	defer resp.Body.Close()
//	// 2.读取页面内容
//	pageBytes, err := ioutil.ReadAll(resp.Body)
//	HandleError(err, "ioutil.ReadAll")
//	// 字节转字符串
//	pageStr := string(pageBytes)
//	//fmt.Println(pageStr)
//	// 3.过滤数据，过滤qq邮箱
//	re := regexp.MustCompile(reQQEmail)
//	// -1代表取全部
//	results := re.FindAllStringSubmatch(pageStr, -1)
//	//fmt.Println(results)
//
//	// 遍历结果
//	for _, result := range results {
//		fmt.Println("email:", result[0])
//		fmt.Println("qq:", result[1])
//	}
//}
//
//// 处理异常
//func HandleError(err error, why string) {
//	if err != nil {
//		fmt.Println(why, err)
//	}
//}
//func main() {
//	GetEmail()
//}

//1.1.2. 正则表达式
//文档：https://studygolang.com/pkgdoc
//API
//re := regexp.MustCompile(reStr)，传入正则表达式，得到正则表达式对象
//ret := re.FindAllStringSubmatch(srcStr,-1)：用正则对象，获取页面页面，srcStr是页面内容，-1代表取全部
//爬邮箱
//方法抽取
//爬超链接
//爬手机号
//http://www.zhaohaowang.com/ 如果连接失效了自己找一个有手机号的就好了
//爬身份证号
//http://henan.qq.com/a/20171107/069413.htm 如果连接失效了自己找一个就好了
//爬图片链接

















