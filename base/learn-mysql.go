// 1、 mysql的使用

// 建库，建表

//CREATE TABLE `person` (
//    `user_id` int(11) NOT NULL AUTO_INCREMENT,
//    `username` varchar(260) DEFAULT NULL,
//    `sex` varchar(260) DEFAULT NULL,
//    `email` varchar(260) DEFAULT NULL,
//    PRIMARY KEY (`user_id`)
//  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
//
//CREATE TABLE place (
//    country varchar(200),
//    city varchar(200),
//    telcode int
//)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

//  mysql> desc person;
//    +----------+--------------+------+-----+---------+----------------+
//    | Field    | Type         | Null | Key | Default | Extra          |
//    +----------+--------------+------+-----+---------+----------------+
//    | user_id  | int(11)      | NO   | PRI | NULL    | auto_increment |
//    | username | varchar(260) | YES  |     | NULL    |                |
//    | sex      | varchar(260) | YES  |     | NULL    |                |
//    | email    | varchar(260) | YES  |     | NULL    |                |
//    +----------+--------------+------+-----+---------+----------------+
//    4 rows in set (0.00 sec)
//
//    mysql> desc place;
//    +---------+--------------+------+-----+---------+-------+
//    | Field   | Type         | Null | Key | Default | Extra |
//    +---------+--------------+------+-----+---------+-------+
//    | country | varchar(200) | YES  |     | NULL    |       |
//    | city    | varchar(200) | YES  |     | NULL    |       |
//    | telcode | int(11)      | YES  |     | NULL    |       |
//    +---------+--------------+------+-----+---------+-------+
//    3 rows in set (0.01 sec)

//1.1.1. mysql使用
//使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动） github.com/jmoiron/sqlx （基于mysql驱动的封装）
//
//命令行输入 ：
//
//    go get github.com/go-sql-driver/mysql
//    go get github.com/jmoiron/sqlx


//链接mysql
//
//    database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
//    //database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")

// ### Insert操作

//package main
//
//import (
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	Db = database
//	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//
//	fmt.Println("insert succ:", id)
//}

// 输出结果

//insert succ: 2

//  ### Select操作

//package main
//
//import (
//	"fmt"
//
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//
//	Db = database
//	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//
//	var person []Person
//	err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//
//	fmt.Println("select succ:", person)
//}

//  输出结果

//  select succ: [{1 stu001 man stu01@qq.com}]

//  ### Update 操作

//package main
//
//import (
//	"fmt"
//
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//
//	Db = database
//	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//
//	res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 1)
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//	row, err := res.RowsAffected()
//	if err != nil {
//		fmt.Println("rows failed, ",err)
//	}
//	fmt.Println("update succ:",row)
//
//}

// 输出结果

//  update succ: 1   //修改后在执行就会是0

//   Delete 操作

//package main
//
//import (
//	"fmt"
//
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//
//	Db = database
//	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
//}
//
//func main() {
//
//	/*
//	   _, err := Db.Exec("delete from person where user_id=?", 1)
//	   if err != nil {
//	       fmt.Println("exec failed, ", err)
//	       return
//	   }
//	*/
//
//	res, err := Db.Exec("delete from person where user_id=?", 1)
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		return
//	}
//
//	row,err := res.RowsAffected()
//	if err != nil {
//		fmt.Println("rows failed, ",err)
//	}
//
//	fmt.Println("delete succ: ",row)
//}

//输出结果

// delete succ:  1

//     ### msyql事物

//1. MySQL事务
//mysql事务特性：
//
//    1) 原子性
//    2) 一致性
//    3) 隔离性
//    4) 持久性

//golang MySQL事务应用：
//
//    1） import (“github.com/jmoiron/sqlx")
//    2)  Db.Begin()        开始事务
//    3)  Db.Commit()        提交事务
//    4)  Db.Rollback()     回滚事务

//package main
//
//import (
//	"fmt"
//
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/jmoiron/sqlx"
//)
//
//type Person struct {
//	UserId   int    `db:"user_id"`
//	Username string `db:"username"`
//	Sex      string `db:"sex"`
//	Email    string `db:"email"`
//}
//
//type Place struct {
//	Country string `db:"country"`
//	City    string `db:"city"`
//	TelCode int    `db:"telcode"`
//}
//
//var Db *sqlx.DB
//
//func init() {
//	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		fmt.Println("open mysql failed,", err)
//		return
//	}
//	Db = database
//}
//
//func main() {
//	conn, err := Db.Begin()
//	if err != nil {
//		fmt.Println("begin failed :", err)
//		return
//	}
//
//	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	fmt.Println("insert succ:", id)
//
//	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	id, err = r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	fmt.Println("insert succ:", id)
//
//	conn.Commit()
//}

// 输出结果

//insert succ: 3
//insert succ: 4

//mysql> select * from person;
//    +---------+----------+------+--------------+
//    | user_id | username | sex  | email        |
//    +---------+----------+------+--------------+
//    |       3 | stu001   | man  | stu01@qq.com |
//    |       4 | stu001   | man  | stu01@qq.com |
//    +---------+----------+------+--------------+
//    2 rows in set (0.00 sec)
















