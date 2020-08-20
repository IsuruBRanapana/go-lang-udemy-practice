package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"udemy/Section-4/todo-api/controller"
	"udemy/Section-4/todo-api/model"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving....")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
