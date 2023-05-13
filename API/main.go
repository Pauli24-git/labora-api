package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"labora-api/API/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetItems).Methods("GET")
	router.HandleFunc("/item", controller.GetItem).Methods("GET")
	router.HandleFunc("/items", controller.CreateItem).Methods("POST")
	router.HandleFunc("/items", controller.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/", controller.DeleteItem).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
