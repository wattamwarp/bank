package controllers

import (
	"net/http"

	"bank.com/bank/app/main/pkg/models"
	"bank.com/bank/app/main/pkg/services"
	"bank.com/bank/app/main/pkg/utils"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	var employess []models.Employees
	employess = services.GetEmployess()

	utils.ReturnResponse(w, employess, http.StatusOK)

}

func CreateEmployees(w http.ResponseWriter, r *http.Request) {

	requestBody := &models.Employees{}
	utils.ParseBody(r, requestBody)

	employeeObject := services.CreateEmployees(requestBody)

	utils.ReturnResponse(w, employeeObject, http.StatusCreated)
}

func UpdateEmployees(w http.ResponseWriter, r *http.Request) {
	requestBody := &models.Employees{}

	utils.ParseBody(r, requestBody)

	employeesObject := services.UpdateEmployees(requestBody)

	utils.ReturnResponse(w, employeesObject, http.StatusOK)
}
