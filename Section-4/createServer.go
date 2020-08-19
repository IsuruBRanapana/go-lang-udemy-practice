package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//end point
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("receive request")
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
