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
	// r.HandleFunc("/api/peternak", GetPeternaks).Methods("GET")
	// r.HandleFunc("/api/peternak/{id}", GetPeternak).Methods("GET")
	r.HandleFunc("/api/login", Login).Methods("POST")
	// r.HandleFunc("/api/peternak/{id}", UpdatePeternak).Methods("PUT")
	// r.HandleFunc("/api/peternak/{id}", DeletePeternak).Methods("DELETE")

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
