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

	var person []Person
	err := Db.Select(&person, "select * from person ")
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	fmt.Println("select succ:", person)
}