package models

import (
	"bank.com/bank/app/main/pkg/config"
	"github.com/jinzhu/gorm"
)

type UserDetails struct {
	gorm.Model
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	PinCode       int       `json:"pincode"`
	PhoneNumber   int       `json:"phonenumber"`
	EmailID       string    `json:"emailId"`
	DOB           string    `json:"dob"`
	IncomeRange   int       `json:"incomeRange"`
	BranchName    string    `json:"branchName"`
	BranchCode    int       `json:"branchCode"`
	BranchAddress string    `json:"branchAddress"`
	EmployeeID    int       `json:"employee_id"` // Foreign key
	Employee      Employees `json:"employees" gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;belongsTo:EmployeeID"`
	IsDeleted     bool      `json:"isDeleted"`
}

var db *gorm.DB

func InitUserDetails() {
	config.DbConnect()
	db = config.GetDbConnection()
	db.AutoMigrate(&UserDetails{})
}
