package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/zafs23/Receipt-Api/receipt-api/middleware"
	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/services"
	"github.com/zafs23/Receipt-Api/receipt-api/storage"

	"github.com/gorilla/mux"
)

type ReceiptHandler struct {
	ShardedStorage *storage.ShardedStorage
	ReceiptCache   map[string]models.IDResponse
}

// var shardedStorage = storage.NewShardedStorage()
// var receiptCache = make(map[string]models.IDResponse)

// handles "/receipts/process" POST endpoint
func (h *ReceiptHandler) ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the validated receipt from the request context
	tempReceipt, ok := r.Context().Value(middleware.ReceiptContextKey).(models.Receipt)
	if !ok {
		http.Error(w, "Invalid receipt data in context", http.StatusInternalServerError)
		return
	}

	// Generate receipt ID and calculate points
	currID := services.GenerateReceiptID(tempReceipt)
	// Check if the ID already exists in the cache
	if response, exists := h.ReceiptCache[currID]; exists {
		// If ID exists, return the cached response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	currPoints, err := services.CalculatePoints(tempReceipt)
	if err != nil {
		http.Error(w, "Error calculating points", http.StatusInternalServerError)
		return
	}

	// Store receipt and points in sharded storage
	h.ShardedStorage.StoreReceipt(currID, tempReceipt, currPoints)

	// Send response with the generated receipt ID
	response := models.IDResponse{ID: currID}
	h.ReceiptCache[currID] = response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handle "/receipts/{id}/points" GET endpoint
func (h *ReceiptHandler) GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the receipt ID from URL variables
	id := mux.Vars(r)["id"]

	// Retrieve receipt and points from the sharded storage
	_, points, found := h.ShardedStorage.GetReceipt(id)
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
