package services

import (
	"bank.com/bank/app/main/pkg/config"
	"bank.com/bank/app/main/pkg/constants"
	"bank.com/bank/app/main/pkg/models"
	"bank.com/bank/app/main/pkg/utils"
)

func GetEmployess() []models.Employees {
	var employees []models.Employees

	error := config.GetDbConnection().Find(&employees).Error

	if error != nil {
		utils.ErrorPrint(constants.Service, error.Error())
	}

	return employees
}

func CreateEmployees(employee *models.Employees) *models.Employees {

	config.GetDbConnection().NewRecord(&employee)

	error := config.GetDbConnection().Create(&employee).Error
	if error != nil {
		utils.ErrorPrint(constants.Service, error.Error())
	}
	return employee
}

func UpdateEmployees(employee *models.Employees) *models.Employees {
	config.GetDbConnection().Update(&employee)

	return employee
}

func deleteEmplyee(employee *models.Employees) *models.Employees {
	employee.IsDeleted = true
	config.GetDbConnection().Update(&employee)

	return employee
}
