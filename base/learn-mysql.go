package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

)

func init()  {
	database, err := sqlx.Open("mysql","root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed",err)
		return
	}
	Db = database
	defer  db.Close()
}

func main()  {
	r, err := Db.Exec("select * from palce")
	if err != nil {
		fmt.Println("select failed",err)
		return
	}
	fmt.Println("seect succ",r)
}
