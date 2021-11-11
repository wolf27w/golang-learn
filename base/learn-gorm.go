//一个神奇的，对开发人员友好的 Golang ORM 库
//
//1.1. 概览
//全特性 ORM (几乎包含所有特性)
//模型关联 (一对一， 一对多，一对多（反向）， 多对多， 多态关联)
//钩子 (Before/After Create/Save/Update/Delete/Find)
//预加载
//事务
//复合主键
//SQL 构造器
//自动迁移
//日志
//基于GORM回调编写可扩展插件
//全特性测试覆盖
//开发者友好

//1.2. 安装
//go get -u github.com/jinzhu/gorm
//1.3. 快速开始
//package main
//
//import (
//  "github.com/jinzhu/gorm"
//  _ "github.com/jinzhu/gorm/dialects/sqlite"
//)
//
//type Product struct {
//  gorm.Model
//  Code string
//  Price uint
//}
//
//func main() {
//  db, err := gorm.Open("sqlite3", "test.db")
//  if err != nil {
//    panic("failed to connect database")
//  }
//  defer db.Close()
//
//  //自动检查 Product 结构是否变化，变化则进行迁移
//  db.AutoMigrate(&Product{})
//
//  // 增
//  db.Create(&Product{Code: "L1212", Price: 1000})
//
//  // 查
//  var product Product
//  db.First(&product, 1) // 找到id为1的产品
//  db.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品
//
//  // 改 - 更新产品的价格为 2000
//  db.Model(&product).Update("Price", 2000)
//
//  // 删 - 删除产品
//  db.Delete(&product)
//}

//1. 模型定义
//1.1. 模型定义
//模型一般都是普通的 Golang 的结构体，Go的基本数据类型，或者指针。sql.Scanner 和 driver.Valuer，同时也支持接口。
//
//例子：
//
//type User struct {
//  gorm.Model
//  Name         string
//  Age          sql.NullInt64
//  Birthday     *time.Time
//  Email        string  `gorm:"type:varchar(100);unique_index"`
//  Role         string  `gorm:"size:255"` //设置字段的大小为255个字节
//  MemberNumber *string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
//  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 Num字段自增
//  Address      string  `gorm:"index:addr"` // 给Address 创建一个名字是  `addr`的索引
//  IgnoreMe     int     `gorm:"-"` //忽略这个字段
//}
//1.2. 结构标签
//标签是声明模型时可选的标记。GORM 支持以下标记：
//
//1.2.1. 支持的结构标签
//标签	说明
//Column	指定列的名称
//Type	指定列的类型
//Size	指定列的大小，默认是 255
//PRIMARY_KEY	指定一个列作为主键
//UNIQUE	指定一个唯一的列
//DEFAULT	指定一个列的默认值
//PRECISION	指定列的数据的精度
//NOT NULL	指定列的数据不为空
//AUTO_INCREMENT	指定一个列的数据是否自增
//INDEX	创建带或不带名称的索引，同名创建复合索引
//UNIQUE_INDEX	类似 索引，创建一个唯一的索引
//EMBEDDED	将 struct 设置为 embedded
//EMBEDDED_PREFIX	设置嵌入式结构的前缀名称
//-	忽略这些字段
//1.2.2. 关联的结构标签
//有关详细信息，请查看「关联」部分
//
//标签	说明
//MANY2MANY	指定连接表名称
//FOREIGNKEY	指定外键
//ASSOCIATION_FOREIGNKEY	指定关联外键
//POLYMORPHIC	指定多态类型
//POLYMORPHIC_VALUE	指定多态的值
//JOINTABLE_FOREIGNKEY	指定连接表的外键
//ASSOCIATION_JOINTABLE_FOREIGNKEY	指定连接表的关联外键
//SAVE_ASSOCIATIONS	是否自动保存关联
//ASSOCIATION_AUTOUPDATE	是否自动更新关联
//ASSOCIATION_AUTOCREATE	是否自动创建关联
//ASSOCIATION_SAVE_REFERENCE	是否引用自动保存的关联
//PRELOAD	是否自动预加载关联













