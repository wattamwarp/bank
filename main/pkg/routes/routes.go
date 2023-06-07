package routes

import (
	"bank.com/bank/main/pkg/controllers"
	"github.com/gorilla/mux"
)

var BankRoutes = func(router *mux.Router) {
	// router.HandleFunc("/usernew", controllers.GetUserNew).Methods("GET")
	router.HandleFunc("/emps", controllers.CreateEmployees).Methods("POST")
	router.HandleFunc("/emps", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user", controllers.Deleteuser).Methods("DELETE")
}
