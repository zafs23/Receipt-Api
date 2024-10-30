package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-api/middleware"
	"receipt-api/models"
	"receipt-api/services"
	"receipt-api/storage"

	"github.com/gorilla/mux"
)

var shardedStorage = storage.NewShardedStorage()

// handles "/receipts/process" POST endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the validated receipt from the request context
	tempReceipt, ok := r.Context().Value(middleware.ReceiptContextKey).(models.Receipt)
	if !ok {
		http.Error(w, "Invalid receipt data in context", http.StatusInternalServerError)
		return
	}

	// Generate receipt ID and calculate points
	currID := services.GenerateReceiptID()
	currPoints, err := services.CalculatePoints(tempReceipt)
	if err != nil {
		http.Error(w, "Error calculating points", http.StatusInternalServerError)
		return
	}

	// Store receipt and points in sharded storage
	shardedStorage.StoreReceipt(currID, tempReceipt, currPoints)

	// Send response with the generated receipt ID
	response := models.IDResponse{ID: currID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handle "/receipts/{id}/points" GET endpoint
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from URL variables
	id := mux.Vars(r)["id"]

	// Retrieve receipt and points from the sharded storage
	_, points, found := shardedStorage.GetReceipt(id)
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Send response with the points
	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	// if err := json.NewEncoder(w).Encode(response); err != nil {
	// 	http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	// 	fmt.Println("Error encoding response:", err)
	// }
}
