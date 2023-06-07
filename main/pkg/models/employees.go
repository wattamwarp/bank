package models

import (
	"bank.com/bank/main/pkg/config"
	"github.com/jinzhu/gorm"
)

type Employees struct {
	gorm.Model
	// Other fields of Employees struct
	Name       string `json:"name"`
	Department string `json:"deparment"`
	Age        int    `json:"age"`
	IsDeleted  bool   `json:"is_deleted"`
}

func EmployeesMigration() {
	config.DbConnect()
	db = config.GetDbConnection()
	db.AutoMigrate(&Employees{})
}
