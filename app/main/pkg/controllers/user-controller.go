package controllers

import (
	"net/http"

	"bank.com/bank/app/main/pkg/models"
	"bank.com/bank/app/main/pkg/services"
	"bank.com/bank/app/main/pkg/utils"
)


func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.UserDetails = services.GetUsers()
	print(users)

	var jsonObject []byte = utils.ReturnJson(users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonObject)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	createUser := &models.UserDetails{}
	utils.ParseBody(r, createUser)
	b := services.CreateUsers(createUser)
	res := utils.ReturnJson(b)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	w.Write(res)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	UpdateUser := &models.UserDetails{}
	utils.ParseBody(r, UpdateUser)
	updatedUser := services.Updateuser(UpdateUser)
	returnJson := utils.ReturnJson(updatedUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write(returnJson)

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	userModel := &models.UserDetails{}
	utils.ParseBody(r, userModel)
	deletedUser := services.DeleteUser(userModel)
	returnJson := utils.ReturnJson(deletedUser)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Context-Type", "pkglication/json")
	w.Write(returnJson)
}
