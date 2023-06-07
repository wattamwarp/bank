package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func DbConnect() {
	d, error := gorm.Open(DbName, DbArguements)

	if error != nil {
		panic(error)
	}
	db = d
}

func GetDbConnection() *gorm.DB {

	return db
}
