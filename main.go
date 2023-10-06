package main

import (
	//"log"

	"log"
	"net/http"
	"time"

	"bank.com/bank/app/main/pkg/models"
	"bank.com/bank/app/main/pkg/routes"
	"github.com/gorilla/mux"
	//"bank.com/bank/app/main/pkg/routes"
	//"github.com/gorilla/mux"
)

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.New()

// 	r.GET("/", func(c *gin.Context) {
// 		fmt.Printf("api is called");
// 		c.String(http.StatusOK, "heloo world this is bank")
// 	})

// 	r.Run(":8080")
// }

func main() {
	println("runnig bank program")

	dbMigration()

	r := mux.NewRouter()

	routes.BankRoutes(r)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	// log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func dbMigration() {
	models.InitUserDetails()
	models.EmployeesMigration()
}
