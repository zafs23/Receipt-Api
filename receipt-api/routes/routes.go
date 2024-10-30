package routes

import (
	"net/http"
	"receipt-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods(http.MethodPost)
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods(http.MethodGet)
	return r
}
