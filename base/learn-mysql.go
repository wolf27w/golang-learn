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

package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}

func main() {
	r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("insert succ:", id)
}