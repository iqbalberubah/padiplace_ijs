package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	//product
	r.HandleFunc("/api/product", GetProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", GetProduct).Methods("GET")
	r.HandleFunc("/api/product", CreateProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", DeleteProduct).Methods("DELETE")

	//peternak
	r.HandleFunc("/api/login", Login).Methods("POST")

	//po
	r.HandleFunc("/api/history", GetHistoryPO).Methods("POST")
	r.HandleFunc("/api/history/detail", GetPO).Methods("POST")

	//balance
	r.HandleFunc("/api/balance", GetBalance).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))

}
