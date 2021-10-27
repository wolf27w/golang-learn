//############Redis介绍

//Redis 简介

//Redis 是完全开源免费的，遵守BSD协议，是一个高性能的key-value数据库。
//
//Redis 与其他 key - value 缓存产品有以下三个特点

//Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
//Redis不仅仅支持简单的key-value类型的数据，同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
//Redis支持数据的备份，即master-slave模式的数据备份。

//    Redis 优势

//性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s ，单机能够达到15w qps，通常适合做缓存。
//
//丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。
//
//原子 – Redis的所有操作都是原子性的，意思就是要么成功执行要么失败完全不执行。单个操作是原子性的。多个操作也支持事务，即原子性，通过MULTI和EXEC指令包起来。
//
//丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。
//
//Redis与其他key-value存储有什么不同？
//
//Redis有着更为复杂的数据结构并且提供对他们的原子性操作，这是一个不同于其他数据库的进化路径。Redis的数据类型都是基于基本数据结构的同时对程序员透明，无需进行额外的抽象。
//
//Redis运行在内存中但是可以持久化到磁盘，所以在对不同数据集进行高速读写时需要权衡内存，因为数据量不能大于硬件内存。在内存数据库方面的另一个优点是，相比在磁盘上相同的复杂的数据结构，在内存中操作起来非常简单，这样Redis可以做很多内部复杂性很强的事情。同时，在磁盘格式方面他们是紧凑的以追加的方式产生的，因为他们并不需要进行随机访问。
//
//1.1.3. redis使用
//使用第三方开源的redis库: github.com/garyburd/redigo/redis
//
//命令行输入 ：
//这里执行是，按照redis模块
//go get github.com/garyburd/redigo/redis

//           连接redis

//package main
//import (
//   "fmt"
//   "github.com/gomodule/redigo/redis"
//)
//func main() {
//   c, err := redis.Dial("tcp", "10.123.6.236:6379")
//   if err != nil {
//       fmt.Println("conn redis failed,", err)
//       return
//   }
//   fmt.Println("redis conn success")
//
//   defer c.Close()
//}

//输出结果：
//redis conn success

// #########redis的set、get操作

//package main
//
//import (
//    "fmt"
//    "github.com/gomodule/redigo/redis"
//)
//
//func main() {
//    c, err := redis.Dial("tcp", "10.123.6.236:6379")
//    if err != nil {
//        fmt.Println("conn redis failed,", err)
//        return
//    }
//
//    defer c.Close()
//    _, err = c.Do("Set", "abc", 100)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    r, err := redis.Int(c.Do("Get", "abc"))
//    if err != nil {
//        fmt.Println("get abc failed,", err)
//        return
//    }
//
//    fmt.Println(r)
//}
//100
//输出结果

//如果出现：    MISCONF Redis is configured to save RDB snapshots, but is currently not able to persist on disk. Commands that may modify the data set are disabled. Please check Redis logs for details about the error.
//Redis被配置为保存数据库快照，但它目前不能持久化到硬盘。用来修改集合数据的命令不能用。请查看Redis日志的详细错误信息。
//原因：
//
//强制关闭Redis快照导致不能持久化。
//
//解决方案：
//
//运行config set stop-writes-on-bgsave-error no　命令后，关闭配置项stop-writes-on-bgsave-error解决该问题。
//在redis上登录后运行127.0.0.1:6379> config set stop-writes-on-bgsave-error n0即可


//  String批量操作

//package main
//
//import (
//    "fmt"
//    "github.com/gomodule/redigo/redis"
//)
//
//func main() {
//    c, err := redis.Dial("tcp", "10.123.6.236:6379")
//    if err != nil {
//        fmt.Println("conn redis failed,", err)
//        return
//    }
//
//    defer c.Close()
//    _, err = c.Do("MSet", "abc", 100, "efg", 300)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
//    if err != nil {
//        fmt.Println("get abc failed,", err)
//        return
//    }
//
//    for _, v := range r {
//        fmt.Println(v)
//    }
//}

//输出结果：
//100
//300
//在redis上操作

//root@e8d032aeb22f:/data# redis-cli -h 127.0.0.1 -p 6379
//127.0.0.1:6379> mget abc efg
//1) "100"
//2) "300"

//   ####设置过期时间

//package main
//
//import (
//    "fmt"
//    "github.com/gomodule/redigo/redis"
//)
//
//func main() {
//    c, err := redis.Dial("tcp", "10.123.6.236:6379")
//    if err != nil {
//        fmt.Println("conn redis failed,", err)
//        return
//    }
//
//    defer c.Close()
//    _, err = c.Do("expire", "abc", 10)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//}

//  Redis命令行窗口：

//   127.0.0.1:6379> get abc
//    "100"
//
//    # 10秒后过期
//    127.0.0.1:6379> get abc
//    (nil)

//       #####List队列操作


//package main
//
//import (
//    "fmt"
//    "github.com/gomodule/redigo/redis"
//)
//
//func main() {
//    c, err := redis.Dial("tcp", "10.123.6.236:6379")
//    if err != nil {
//        fmt.Println("conn redis failed,", err)
//        return
//    }
//
//    defer c.Close()
//    _, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
//    if err != nil {
//        fmt.Println(err)
//        return
//    }
//
//    r, err := redis.String(c.Do("lpop", "book_list"))
//    if err != nil {
//        fmt.Println("get abc failed,", err)
//        return
//    }
//
//    fmt.Println(r)
//}

//输出结果
// 300

//redis 命令行

// 127.0.0.1:6379> lpop book_list
//    "ceg"
//    127.0.0.1:6379> lpop book_list
//    "abc"
//    127.0.0.1:6379> lpop book_list
//    (nil)


//    #####Hash表

package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "10.123.6.236:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }

    defer c.Close()
    _, err = c.Do("HSet", "books", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }

    r, err := redis.Int(c.Do("HGet", "books", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }

    fmt.Println(r)
}












