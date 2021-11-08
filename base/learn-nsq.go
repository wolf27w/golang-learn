//  安装
//Mac安装nsq：
//
//按照安装文档中的说明进行操作。
//
//打开终端：
//
//执行：$ brew install nsq
//
//若：-bash: brew: command not found
//
//执行：$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
//
//然后执行：brew install nsq
//
//在一个shell中，开始nsqlookupd：
//
//$ nsqlookupd
//
//在另一个shell中，开始nsqd：
//
//$ nsqd --lookupd-tcp-address=127.0.0.1:4160
//
//在另一个shell中，开始nsqadmin：
//
//$ nsqadmin --lookupd-http-address=127.0.0.1:4161
//
//发布初始消息（也在集群中创建主题）：
//
//$ curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test'
//
//最后，在另一个shell中，开始nsq_to_file：
//
//$ nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
//
//发布更多消息nsqd：
//
//$ curl -d 'hello world 2' 'http://127.0.0.1:4151/pub?topic=test'
//
//$ curl -d 'hello world 3' 'http://127.0.0.1:4151/pub?topic=test'
//
//验证事物按预期工作，在Web浏览器中打开http://127.0.0.1:4171/ 以查看nsqadminUI并查看统计信息。另外，检查test.*.log写入的日志文件（）的内容/tmp。
//
//这里的重要教训是nsq_to_file（客户端）未明确告知test 主题产生的位置，它从中检索此信息，nsqlookupd并且尽管有连接的时间，但不会丢失任何消息。

//生产者
//运行Nsq服务集群
//
//首先启动nsqlookud，在一个shell中，开始nsqlookupd：
//
//$ nsqlookupd
//
//在另一个shell中，开始nsqd：
//
//$ nsqd --lookupd-tcp-address=127.0.0.1:4160
//
//在另一个shell中，开始nsqadmin：
//
//$ nsqadmin --lookupd-http-address=127.0.0.1:4161
//
//发布初始消息（也在集群中创建主题）：
//
//$ curl -d 'hello world 1' 'http://127.0.0.1:4151/pub?topic=test'
//
//最后，在另一个shell中，开始nsq_to_file：
//
//$ nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
//
//验证事物按预期工作，在Web浏览器中打开http://127.0.0.1:4171/ 以查看nsqadminUI并查看统计信息。另外，检查test.*.log写入的日志文件（）的内容/tmp。
//
//链接nsq 并创建生产者：

//package main
//
//import (
//    "fmt"
//
//    nsq "github.com/nsqio/go-nsq"
//)
//
//func main() {
//    // 定义nsq生产者
//    var producer *nsq.Producer
//    // 初始化生产者
//    // producer, err := nsq.NewProducer("地址:端口", nsq.*Config )
//    producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
//    if err != nil {
//        panic(err)
//    }
//
//    err = producer.Ping()
//    if nil != err {
//        // 关闭生产者
//        producer.Stop()
//        producer = nil
//    }
//
//    fmt.Println("ping nsq success")
//}

//生产者创建topic并写入nsq：