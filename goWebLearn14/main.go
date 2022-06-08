package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/springbootwebtest")

	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()

	err = db.PingContext(ctx)

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected")
	//查询
	one, err := getOne(1)

	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(one)
	one.role += "1"
	err = one.Update()
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(one)
	one2, err := getMany(0)

	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(one2)

	//添加有问题
	t := teacher{
		username: "嘉然",
		password: "123456",
		role:     "3",
	}
	err = t.Insert()
	if err != nil {
		log.Fatalln(err.Error())
	}
	one3, _ := getOne(t.tNo)
	fmt.Println(one3)
}
