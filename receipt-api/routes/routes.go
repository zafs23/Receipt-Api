package routes

import (
	"net/http"
	"receipt-api/handlers"
	"receipt-api/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/receipts/process", middleware.ValidateReceiptMiddleware(http.HandlerFunc(handlers.ProcessReceiptHandler))).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	return r
}
