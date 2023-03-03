package handler

import (
	"encoding/json"
	"net/http"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
)

func CreateTransactionInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var transactionDetail e.TransactionDetail
	json.NewDecoder(r.Body).Decode(&transactionDetail)
	d.DB.Table("tbl_invoice").Create(&transactionDetail)
	json.NewEncoder(w).Encode(transactionDetail)
}
