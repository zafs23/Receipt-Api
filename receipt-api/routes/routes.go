package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zafs23/Receipt-Api/receipt-api/handlers"
	"github.com/zafs23/Receipt-Api/receipt-api/middleware"
)

func RegisterRoutes(receiptHandler *handlers.ReceiptHandler) *mux.Router {
	r := mux.NewRouter()
	// r.Handle("/receipts/process", middleware.ValidateReceiptMiddleware(http.HandlerFunc(handlers.ProcessReceiptHandler))).Methods("POST")
	// r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	// Use the struct methods directly on the receiptHandler instance
	r.Handle("/receipts/process", middleware.ValidateReceiptMiddleware(http.HandlerFunc(receiptHandler.ProcessReceiptHandler))).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", receiptHandler.GetPointsHandler).Methods("GET")

	return r
}
