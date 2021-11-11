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

//               grud 接口

//1. 创建
//1.1. 创建记录
//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
//
//db.NewRecord(user) // => 返回 `true` ，因为主键为空
//
//db.Create(&user)
//
//db.NewRecord(user) // => 在 `user` 之后创建返回 `false`
//1.2. 默认值
//你可以通过标签定义字段的默认值，例如：
//
//type Animal struct {
//    ID   int64
//    Name string `gorm:"default:'galeone'"`
//    Age  int64
//}
//然后 SQL 会排除那些没有值或者有 零值 的字段，在记录插入数据库之后，gorm将从数据库中加载这些字段的值。
//
//var animal = Animal{Age: 99, Name: ""}
//db.Create(&animal)
//// INSERT INTO animals("age") values('99');
//// SELECT name from animals WHERE ID=111; // 返回的主键是 111
//// animal.Name => 'galeone'
//注意 所有包含零值的字段，像 0，''，false 或者其他的 零值 不会被保存到数据库中，但会使用这个字段的默认值。你应该考虑使用指针类型或者其他的值来避免这种情况:
//
//// Use pointer value
//type User struct {
//  gorm.Model
//  Name string
//  Age  *int `gorm:"default:18"`
//}
//
//// Use scanner/valuer
//type User struct {
//  gorm.Model
//  Name string
//  Age  sql.NullInt64 `gorm:"default:18"`
//}
//1.3. 在钩子中设置字段值
//如果你想在 BeforeCreate 函数中更新字段的值，应该使用 scope.SetColumn，例如：
//
//func (user *User) BeforeCreate(scope *gorm.Scope) error {
//  scope.SetColumn("ID", uuid.New())
//  return nil
//}
//1.4. 创建额外选项
//// 为插入 SQL 语句添加额外选项
//db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)
//// INSERT INTO products (name, code) VALUES ("name", "code") ON CONFLICT;

//1. 查询
//1.1. 查询
//// 获取第一条记录，按主键排序
//db.First(&user)
////// SELECT * FROM users ORDER BY id LIMIT 1;
//
//// 获取一条记录，不指定排序
//db.Take(&user)
////// SELECT * FROM users LIMIT 1;
//
//// 获取最后一条记录，按主键排序
//db.Last(&user)
////// SELECT * FROM users ORDER BY id DESC LIMIT 1;
//
//// 获取所有的记录
//db.Find(&users)
////// SELECT * FROM users;
//
//// 通过主键进行查询 (仅适用于主键是数字类型)
//db.First(&user, 10)
////// SELECT * FROM users WHERE id = 10;
//1.1.1. Where
//原生 SQL
//// 获取第一条匹配的记录
//db.Where("name = ?", "jinzhu").First(&user)
////// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
//
//// 获取所有匹配的记录
//db.Where("name = ?", "jinzhu").Find(&users)
////// SELECT * FROM users WHERE name = 'jinzhu';
//
//// <>
//db.Where("name <> ?", "jinzhu").Find(&users)
//
//// IN
//db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
//
//// LIKE
//db.Where("name LIKE ?", "%jin%").Find(&users)
//
//// AND
//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
//
//// Time
//db.Where("updated_at > ?", lastWeek).Find(&users)
//
//// BETWEEN
//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
//Struct & Map
//// Struct
//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
////// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
//
//// Map
//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
////// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;
//
//// 多主键 slice 查询
//db.Where([]int64{20, 21, 22}).Find(&users)
////// SELECT * FROM users WHERE id IN (20, 21, 22);
//NOTE 当通过struct进行查询的时候，GORM 将会查询这些字段的非零值， 意味着你的字段包含 0， ''， false 或者其他 零值, 将不会出现在查询语句中， 例如:
//
//db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
////// SELECT * FROM users WHERE name = "jinzhu";
//你可以考虑适用指针类型或者 scanner/valuer 来避免这种情况。
//
//// 使用指针类型
//type User struct {
//  gorm.Model
//  Name string
//  Age  *int
//}
//
//// 使用 scanner/valuer
//type User struct {
//  gorm.Model
//  Name string
//  Age  sql.NullInt64
//}
//1.1.2. Not
//和 Where查询类似
//
//db.Not("name", "jinzhu").First(&user)
////// SELECT * FROM users WHERE name <> "jinzhu" LIMIT 1;
//
//// 不包含
//db.Not("name", []string{"jinzhu", "jinzhu 2"}).Find(&users)
////// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//
////不在主键 slice 中
//db.Not([]int64{1,2,3}).First(&user)
////// SELECT * FROM users WHERE id NOT IN (1,2,3);
//
//db.Not([]int64{}).First(&user)
////// SELECT * FROM users;
//
//// 原生 SQL
//db.Not("name = ?", "jinzhu").First(&user)
////// SELECT * FROM users WHERE NOT(name = "jinzhu");
//
//// Struct
//db.Not(User{Name: "jinzhu"}).First(&user)
////// SELECT * FROM users WHERE name <> "jinzhu";
//1.1.3. Or
//db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
////// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
//
//// Struct
//db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2"}).Find(&users)
////// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';
//
//// Map
//db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2"}).Find(&users)
////// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';


//1. 更新
//1.1. 更新所有字段
//Save 方法在执行 SQL 更新操作时将包含所有字段，即使这些字段没有被修改。
//
//db.First(&user)
//
//user.Name = "jinzhu 2"
//user.Age = 100
//db.Save(&user)
//
////// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
//1.2. 更新已更改的字段
//如果你只想更新已经修改了的字段，可以使用 Update，Updates 方法。
//
//// 如果单个属性被更改了，更新它
//db.Model(&user).Update("name", "hello")
////// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
//
//// 使用组合条件更新单个属性
//db.Model(&user).Where("active = ?", true).Update("name", "hello")
////// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
//
//// 使用 `map` 更新多个属性，只会更新那些被更改了的字段
//db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
////// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
//
//// 使用 `struct` 更新多个属性，只会更新那些被修改了的和非空的字段
//db.Model(&user).Updates(User{Name: "hello", Age: 18})
////// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
//
//// 警告： 当使用结构体更新的时候, GORM 只会更新那些非空的字段
//// 例如下面的更新，没有东西会被更新，因为像 "", 0, false 是这些字段类型的空值
//db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})
//1.3. 更新选中的字段
//如果你在执行更新操作时只想更新或者忽略某些字段，可以使用 Select，Omit方法。
//
//db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
////// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
//
//db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
////// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
//1.4. 更新列钩子方法
//上面的更新操作更新时会执行模型的 BeforeUpdate 和 AfterUpdate 方法，来更新 UpdatedAt 时间戳，并且保存他的 关联。如果你不想执行这些操作，可以使用 UpdateColumn，UpdateColumns 方法。
//
//// Update single attribute, similar with `Update`
//db.Model(&user).UpdateColumn("name", "hello")
////// UPDATE users SET name='hello' WHERE id = 111;
//
//// Update multiple attributes, similar with `Updates`
//db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
////// UPDATE users SET name='hello', age=18 WHERE id = 111;
//1.5. 批量更新
//批量更新时，钩子函数不会执行
//
//db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
////// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
//
//// 使用结构体更新将只适用于非零值，或者使用 map[string]interface{}
//db.Model(User{}).Updates(User{Name: "hello", Age: 18})
////// UPDATE users SET name='hello', age=18;
//
//// 使用 `RowsAffected` 获取更新影响的记录数
//db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected
//1.6. 带有表达式的 SQL 更新
//DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
////// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';
//
//DB.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
////// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';
//
//DB.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
////// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';
//
//DB.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
////// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;
//1.7. 在钩子函数中更新值
//如果你想使用 BeforeUpdate、BeforeSave钩子函数修改更新的值，可以使用 scope.SetColumn方法，例如：
//
//func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
//  if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
//    scope.SetColumn("EncryptedPassword", pw)
//  }
//}
//1.8. 额外的更新选项
//// 在更新 SQL 语句中添加额外的 SQL 选项
//db.Model(&user).Set("gorm:update_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Update("name", "hello")
////// UPDATE users SET name='hello', updated_at = '2013-11-17 21:34:10' WHERE id=111 OPTION (OPTIMIZE FOR UNKNOWN);

//1. 删除
//1.1. 删除记录
//警告：当删除一条记录的时候，你需要确定这条记录的主键有值，GORM会使用主键来删除这条记录。如果主键字段为空，GORM会删除模型中所有的记录。
//
//// 删除一条存在的记录
//db.Delete(&email)
////// DELETE from emails where id=10;
//
//// 为删除 SQL 语句添加额外选项
//db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&email)
////// DELETE from emails where id=10 OPTION (OPTIMIZE FOR UNKNOWN);
//1.2. 批量删除
//删除所有匹配的记录
//
//db.Where("email LIKE ?", "%jinzhu%").Delete(Email{})
////// DELETE from emails where email LIKE "%jinzhu%";
//
//db.Delete(Email{}, "email LIKE ?", "%jinzhu%")
////// DELETE from emails where email LIKE "%jinzhu%";
//1.3. 软删除
//如果模型中有 DeletedAt 字段，它将自动拥有软删除的能力！当执行删除操作时，数据并不会永久的从数据库中删除，而是将 DeletedAt 的值更新为当前时间。
//
//db.Delete(&user)
////// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;
//
//// 批量删除
//db.Where("age = ?", 20).Delete(&User{})
////// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
//
//// 在查询记录时，软删除记录会被忽略
//db.Where("age = 20").Find(&user)
////// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
//
//// 使用 Unscoped 方法查找软删除记录
//db.Unscoped().Where("age = 20").Find(&users)
////// SELECT * FROM users WHERE age = 20;
//
//// 使用 Unscoped 方法永久删除记录
//db.Unscoped().Delete(&order)
////// DELETE FROM orders WHERE id=10;

//        关联

//1. Belongs To
//1.1. 属于
//belongs to 关联建立一个和另一个模型的一对一连接，使得模型声明每个实例都「属于」另一个模型的一个实例 。
//
//例如，如果你的应用包含了用户和用户资料， 并且每一个用户资料只分配给一个用户
//
//type User struct {
//  gorm.Model
//  Name string
//}
//
//// `Profile` 属于 `User`， `UserID` 是外键
//type Profile struct {
//  gorm.Model
//  UserID int
//  User   User
//  Name   string
//}
//1.2. 外键
//为了定义从属关系， 外键是必须存在的， 默认的外键使用所有者类型名称加上其主键。
//
//像上面的例子，为了声明一个模型属于 User，它的外键应该为 UserID。
//
//GORM 提供了一个定制外键的方法，例如:
//
//type User struct {
//    gorm.Model
//    Name string
//}
//
//type Profile struct {
//    gorm.Model
//  Name      string
//  User      User `gorm:"foreignkey:UserRefer"` // 使用 UserRefer 作为外键
//  UserRefer string
//}
//1.3. 关联外键
//对于从属关系， GORM 通常使用所有者的主键作为外键值，在上面的例子中，就是 User 的 ID。
//
//当你分配一个资料给一个用户， GORM 将保存用户表的 ID 值 到 用户资料表的 UserID 字段里。
//
//你可以通过改变标签 association_foreignkey 来改变它， 例如：
//
//type User struct {
//    gorm.Model
//  Refer int
//    Name string
//}
//
//type Profile struct {
//    gorm.Model
//  Name      string
//  User      User `gorm:"association_foreignkey:Refer"` // use Refer 作为关联外键
//  UserRefer string
//}
//1.4. 使用属于
//你能找到 belongs to 和 Related 的关联
//
//db.Model(&user).Related(&profile)
////// SELECT * FROM profiles WHERE user_id = 111; // 111 is user's ID
//更多高级用法，请参考 Association Mode


//1. Has One
//1.1. Has One
//has one 关联也是与另一个模型建立一对一的连接，但语义（和结果）有些不同。 此关联表示模型的每个实例包含或拥有另一个模型的一个实例。
//
//例如，如果你的应用程序包含用户和信用卡，并且每个用户只能有一张信用卡。
//
//// 用户有一个信用卡，CredtCardID 外键
//type User struct {
//    gorm.Model
//    CreditCard   CreditCard
//}
//
//type CreditCard struct {
//    gorm.Model
//    Number   string
//    UserID    uint
//}
//1.2. 外键
//对于一对一关系，一个外键字段也必须存在，所有者将保存主键到模型关联的字段里。
//
//这个字段的名字通常由 belongs to model 的类型加上它的 primary key 产生的，就上面的例子而言，它就是 CreditCardID
//
//当你给用户一个信用卡， 它将保存一个信用卡的 ID 到 CreditCardID 字段中。
//
//如果你想使用另一个字段来保存这个关系，你可以通过使用标签 foreignkey 来改变它， 例如：
//
//type User struct {
//  gorm.Model
//  CreditCard CreditCard `gorm:"foreignkey:CardRefer"`
//}
//
//type CreditCard struct {
//    gorm.Model
//    Number      string
//    UserName string
//}
//1.3. 关联外键
//通常，所有者会保存 belogns to model 的主键到外键，你可以改为保存其他字段， 就像下面的例子使用 Number 。
//
//type User struct {
//  gorm.Model
//  CreditCard CreditCard `gorm:"association_foreignkey:Number"`
//}
//
//type CreditCard struct {
//    gorm.Model
//    Number string
//    UID       string
//}
//1.4. 多态关联
//支持多态的一对多和一对一关联。
//
//  type Cat struct {
//    ID    int
//    Name  string
//    Toy   Toy `gorm:"polymorphic:Owner;"`
//  }
//
//  type Dog struct {
//    ID   int
//    Name string
//    Toy  Toy `gorm:"polymorphic:Owner;"`
//  }
//
//  type Toy struct {
//    ID        int
//    Name      string
//    OwnerID   int
//    OwnerType string
//  }
//注意：多态属于和多对多是明确的不支持并将会抛出错误。
//
//1.5. 使用一对一
//你可以通过 Related 找到 has one 关联。
//
//var card CreditCard
//db.Model(&user).Related(&card, "CreditCard")
////// SELECT * FROM credit_cards WHERE user_id = 123; // 123 是用户表的主键
//// CreditCard  是用户表的字段名，这意味着获取用户的信用卡关系并写入变量 card。
//// 像上面的例子，如果字段名和变量类型名一样，它就可以省略， 像：
//db.Model(&user).Related(&card)
//更多高级用法，请参考 Association Mode

