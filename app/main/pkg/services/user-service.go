package services

import (
	"bank.com/bank/app/main/pkg/config"
	"bank.com/bank/app/main/pkg/models"
	"bank.com/bank/app/main/pkg/utils"
)


func GetUsers() []models.UserDetails {
	var users []models.UserDetails

	error :=
		config.GetDbConnection().Table("user_details").Preload("Employee").Find(&users).Error

	if error != nil {
		utils.ErrorPrint("Service", error.Error())
	}
	return users
}

func CreateUsers(userModel *models.UserDetails) *models.UserDetails {
	config.GetDbConnection().Table("user_details").NewRecord(&userModel)
	config.GetDbConnection().Create(&userModel)
	return userModel
}

func Updateuser(userModel *models.UserDetails) *models.UserDetails {
	config.GetDbConnection().Model(&userModel).Updates(userModel)
	return userModel
}

func DeleteUser(usermodel *models.UserDetails) *models.UserDetails {
	usermodel.IsDeleted = true
	print("*********** userModel is", usermodel)
	config.GetDbConnection().Model(&usermodel).Updates(usermodel)
	return usermodel
}
