package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) //untuk menampilkan status
		w.Write([]byte("Hello World"))
	}).Methods("GET")

	fmt.Println("server running on port 2000")
	http.ListenAndServe("localhost:2000", route)
}
