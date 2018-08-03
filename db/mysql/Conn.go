package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123@(127.0.0.1:3306)/gotest?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	//defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
