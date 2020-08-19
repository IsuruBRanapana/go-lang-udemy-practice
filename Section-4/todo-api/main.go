package main

import (
	"net/http"
	"udemy/Section-4/todo-api/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
