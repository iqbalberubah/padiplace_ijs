package router

import (
	"log"
	"net/http"
	h "padiplace_ijs/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()
	//product
	r.HandleFunc("/api/product", h.GetProducts).Methods("GET")
	r.HandleFunc("/api/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/api/product", h.CreateProduct).Methods("POST")
	r.HandleFunc("/api/product/{id}", h.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", h.DeleteProduct).Methods("DELETE")

	//peternak
	r.HandleFunc("/api/login", h.Login).Methods("POST")

	//po
	r.HandleFunc("/api/history", h.GetHistoryPO).Methods("POST")
	r.HandleFunc("/api/history/detail", h.GetPO).Methods("POST")

	//balance
	r.HandleFunc("/api/balance", h.GetBalance).Methods("POST")

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
