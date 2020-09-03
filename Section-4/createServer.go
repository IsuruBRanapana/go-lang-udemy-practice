package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//end point
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receive request")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World"))
	})

	mux.HandleFunc("/getgoing/go", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receive request")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World go"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
