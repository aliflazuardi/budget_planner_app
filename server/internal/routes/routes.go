package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	// create new mux router
	r := http.NewServeMux()

	// register routes
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/api/data", apidataHandler)
	// r.HandleFunc("/api/transactions", getTransactions)
	// r.HandleFunc("/api/transactions", createTransaction)
	// r.HandleFunc("/api/transactions/{id}", updateTransaction)
	// r.HandleFunc("/api/transactions/{id}", deleteTransaction)

	// return router
	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to  Budget Planner App"))
}

func apidataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Data"))
}
