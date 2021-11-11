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

//1. 惯例
//1.1. gorm.Model
//gorm.Model 是一个包含一些基本字段的结构体, 包含的字段有 ID，CreatedAt， UpdatedAt， DeletedAt。
//
//你可以用它来嵌入到你的模型中，或者也可以用它来建立自己的模型。
//
//// gorm.Model 定义
//type Model struct {
//  ID        uint `gorm:"primary_key"`
//  CreatedAt time.Time
//  UpdatedAt time.Time
//  DeletedAt *time.Time
//}
//
//// 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入到 `User` 模型中
//type User struct {
//  gorm.Model
//  Name string
//}
//
//// 声明 gorm.Model 模型
//type User struct {
//  ID   int
//  Name string
//}
//1.2. ID 作为主键
//GORM 默认使用 ID 作为主键名。
//
//type User struct {
//  ID   string // 字段名 `ID` 将被作为默认的主键名
//}
//
//// 设置字段 `AnimalID` 为默认主键
//type Animal struct {
//  AnimalID int64 `gorm:"primary_key"`
//  Name     string
//  Age      int64
//}
//1.3. 复数表名
//表名是结构体名称的复数形式
//
//type User struct {} // 默认的表名是 `users`
//
//// 设置 `User` 的表名为 `profiles`
//func (User) TableName() string {
//  return "profiles"
//}
//
//func (u User) TableName() string {
//    if u.Role == "admin" {
//        return "admin_users"
//    } else {
//        return "users"
//    }
//}
//
//// 如果设置禁用表名复数形式属性为 true，`User` 的表名将是 `user`
//db.SingularTable(true)
//1.3.1. 指定表名
//// 用 `User` 结构体创建 `delete_users` 表
//db.Table("deleted_users").CreateTable(&User{})
//
//var deleted_users []User
//db.Table("deleted_users").Find(&deleted_users)
////// SELECT * FROM deleted_users;
//
//db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
////// DELETE FROM deleted_users WHERE name = 'jinzhu';
//1.3.2. 修改默认表名
//你可以通过定义 DefaultTableNameHandler 字段来对表名使用任何规则。
//
//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
//    return "prefix_" + defaultTableName;
//}
//1.4. 蛇形列名
//列名是字段名的蛇形小写形式
//
//type User struct {
//  ID        uint      // 字段名是 `id`
//  Name      string    // 字段名是 `name`
//  Birthday  time.Time // 字段名是 `birthday`
//  CreatedAt time.Time // 字段名是 `created_at`
//}
//
//// 重写列名
//type Animal struct {
//    AnimalId    int64     `gorm:"column:beast_id"`         // 设置列名为 `beast_id`
//    Birthday    time.Time `gorm:"column:day_of_the_beast"` // 设置列名为 `day_of_the_beast`
//    Age         int64     `gorm:"column:age_of_the_beast"` // 设置列名为 `age_of_the_beast`
//}
//1.5. 时间戳跟踪
//1.5.1. CreatedAt
//对于有 CreatedAt 字段的模型，它将被设置为首次创建记录的当前时间。
//
//db.Create(&user) // 将设置 `CreatedAt` 为当前时间
//
//// 你可以使用 `Update` 方法来更改默认时间
//db.Model(&user).Update("CreatedAt", time.Now())
//1.5.2. UpdatedAt
//对于有 UpdatedAt 字段的模型，它将被设置为记录更新时的当前时间。
//
//db.Save(&user) // 将设置 `UpdatedAt` 为当前时间
//
//db.Model(&user).Update("name", "jinzhu") // 将设置 `UpdatedAt` 为当前时间
//1.5.3. DeletedAt
//对于有 DeletedAt 字段的模型，当删除它们的实例时，它们并没有被从数据库中删除，只是将 DeletedAt 字段设置为当前时间。参考 Soft Delete

//1. 连接数据库
//1.1. 连接数据库
//为了连接数据库，你首先要导入数据库驱动程序。例如：
//
//import _ "github.com/go-sql-driver/mysql"
//GORM 已经包含了一些驱动程序，为了方便的去记住它们的导入路径，你可以像下面这样导入 mysql 驱动程序
//
//import _ "github.com/jinzhu/gorm/dialects/mysql"
//// import _ "github.com/jinzhu/gorm/dialects/postgres"
//// import _ "github.com/jinzhu/gorm/dialects/sqlite"
//// import _ "github.com/jinzhu/gorm/dialects/mssql"
//1.2. 支持的数据库
//1.2.1. MySQL
//注意： 为了正确的处理 time.Time ，你需要包含 parseTime 作为参数。 (More supported parameters)
//
//import (
//  "github.com/jinzhu/gorm"
//  _ "github.com/jinzhu/gorm/dialects/mysql"
//)
//
//func main() {
//  db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
//  defer db.Close()
//}
//1.2.2. PostgreSQL
//import (
//  "github.com/jinzhu/gorm"
//  _ "github.com/jinzhu/gorm/dialects/postgres"
//)
//
//func main() {
//  db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
//  defer db.Close()
//}
//1.2.3. Sqlite3
//import (
//  "github.com/jinzhu/gorm"
//  _ "github.com/jinzhu/gorm/dialects/sqlite"
//)
//
//func main() {
//  db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
//  defer db.Close()
//}
//1.2.4. SQL Server
//Get started with SQL Server，它可以通过 Docker 运行在你的 Mac， Linux 上。
//
//import (
//  "github.com/jinzhu/gorm"
//  _ "github.com/jinzhu/gorm/dialects/mssql"
//)
//
//func main() {
//  db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
//  defer db.Close()
//}
//1.3. 不支持的数据库
//GORM 官方支持以上四种数据库， 你可以为不支持的数据库编写支持，参考 GORM Dialects









