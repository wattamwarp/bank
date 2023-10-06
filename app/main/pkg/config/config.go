package config

import (
	"fmt"
	"os"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func DbConnect() {
	fmt.Printf("%s\n", os.Getenv("DB"))
	// DBMS := os.Getenv("DB")
	// USER := os.Getenv("MYSQL_USER")
	// PASS := os.Getenv("MYSQL_PASSWORD")
	// PROTOCOL := os.Getenv("PROTOCOL")
	// DBNAME := os.Getenv("MYSQL_DATABASE")

	// CONNECT := USER + ":" + PASS + "@tcp(" + PROTOCOL + ":3306)" + "/" + DBNAME + "?parseTime=true"
	d, error := gorm.Open(DbName, DbArguements)
	if error != nil {
		panic(error.Error())
	}
	db = d
}

func GetDbConnection() *gorm.DB {
	return db
}
