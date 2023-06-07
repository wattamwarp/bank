package main

import (
	"log"
	"net/http"

	"bank.com/bank/main/pkg/models"
	"bank.com/bank/main/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	println("runnig bank program")

	dbMigration()

	r := mux.NewRouter()

	routes.BankRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8082", r))

}

func dbMigration() {
	models.InitUserDetails()
	models.EmployeesMigration()
	models.CompanyMigration()
}
