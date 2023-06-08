package models

import (
	"bank.com/bank/app/main/pkg/config"
	"github.com/jinzhu/gorm"
)

// type UserDetails struct {
// 	gorm.Model
// 	Name          string    `json:"name"`
// 	Address       string    `json:"address"`
// 	PinCode       int       `json:"pincode"`
// 	PhoneNumber   int       `json:"phonenumber"`
// 	EmailID       string    `json:"emailId"`
// 	DOB           string    `json:"dob"`
// 	IncomeRange   int       `json:"incomeRange"`
// 	BranchName    string    `json:"branchName"`
// 	BranchCode    int       `json:"branchCode"`
// 	BranchAddress string    `json:"branchAddress"`
// 	EmployeeID    uint64    `json:"employee_id"` // Foreign key
// 	Employee      Employees `json:"employees" gorm:"foreignKey:EmployeeID;belongsTo:EmployeeID"`
// 	IsDeleted     bool      `json:"isDeleted"`
// }

// type UserDetails struct {
// 	gorm.Model
// 	Name          string    `json:"name"`
// 	Address       string    `json:"address"`
// 	PinCode       int       `json:"pincode"`
// 	PhoneNumber   int       `json:"phonenumber"`
// 	EmailID       string    `json:"emailId"`
// 	DOB           string    `json:"dob"`
// 	IncomeRange   int       `json:"incomeRange"`
// 	BranchName    string    `json:"branchName"`
// 	BranchCode    int       `json:"branchCode"`
// 	BranchAddress string    `json:"branchAddress"`
// 	EmployeeID    uint64    `json:"employee_id" gorm:"foreignKey:employee_id"`
// 	Employee      Employees `json:"employees"`
// 	IsDeleted     bool      `json:"isDeleted"`
// }

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

// func (u *UserDetails) CreateUsers() *UserDetails {
// 	config.GetDbConnection().NewRecord(&u)
// 	config.GetDbConnection().Create(&u)
// 	return u
// }

// func (user *UserDetails) updateUserDetails() *UserDetails {

// 	config.GetDbConnection().Save(&user)

// 	return user

// }
