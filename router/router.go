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

	//stock
	// r.HandleFunc("/api/stock", h.GetStocks).Methods("GET")
	// r.HandleFunc("/api/stock/{id}", h.GetStock).Methods("GET")
	// r.HandleFunc("/api/stock", h.CreateStock).Methods("POST")
	// r.HandleFunc("/api/stock/{id}", h.UpdateStock).Methods("PUT")
	// r.HandleFunc("/api/stock/{id}", h.DeleteStock).Methods("DELETE")

	//file
	r.HandleFunc("/api/file", h.GetFiles).Methods("GET")
	r.HandleFunc("/api/file/{id}", h.GetFile).Methods("GET")
	r.HandleFunc("/api/file", h.CreateFile).Methods("POST")
	r.HandleFunc("/api/file/{id}", h.UpdateFile).Methods("PUT")
	r.HandleFunc("/api/file/{id}", h.DeleteFile).Methods("DELETE")

	//kios
	r.HandleFunc("/api/login", h.Login).Methods("POST")
	r.HandleFunc("/api/register", h.Register).Methods("POST")
	r.HandleFunc("/api/user/{id}", h.UpdateTokenFcm).Methods("PUT")
	r.HandleFunc("/api/user", h.GetUsers).Methods("GET")

	//po
	r.HandleFunc("/api/transaction", h.TransactionHistory).Methods("POST")
	r.HandleFunc("/api/transaction/all", h.GetAll).Methods("GET")
	r.HandleFunc("/api/transaction/detail", h.TransactionDetail).Methods("POST")
	r.HandleFunc("/api/transaction/create", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/api/transaction/{id}", h.UpdateTransaction).Methods("PUT")

	//balance
	r.HandleFunc("/api/balance", h.GetBalance).Methods("POST")

	//notification
	// r.HandleFunc("/api/notification", h.GetNotification).Methods("POST")

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
