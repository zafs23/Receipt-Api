package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"receipt-api/models"
	"receipt-api/services"
	"sync"

	"github.com/gorilla/mux"
)

var (
	receiptsMap = make(map[string]models.Receipt)
	pointsMap   = make(map[string]int)
	mutex       = sync.Mutex{}
)

// handles "/receipts/process" POST endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var curr_receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&curr_receipt); err != nil {
		// handle the error and send 400 code
		http.Error(w, "Invalid json request", http.StatusBadRequest)
		return
	}

	// generate the ID
	curr_id := services.GenerateReceiptID()
	curr_receipt_points, err := services.CalculatePoints(curr_receipt)
	if err != nil {
		// Handle the error returned from CalculatePoints
		http.Error(w, fmt.Sprintf("Error calculating points: %v", err), http.StatusInternalServerError)
		return
	}

	// lock the current receipt while updating the value so no other go routine can access it
	mutex.Lock()
	receiptsMap[curr_id] = curr_receipt
	pointsMap[curr_id] = curr_receipt_points
	mutex.Unlock()
	// unlock

	response := models.IDResponse{ID: curr_id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// handle "/receipts/{id}/points" GET endpoint
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	curr_id := mux.Vars(r)["id"] // get the {id} from endpoint

	// lock the current receipt while updating the value so no other go routine can access it
	mutex.Lock()
	curr_points, ok := pointsMap[curr_id]
	mutex.Unlock()

	if !ok {
		// handle the error and send 404 code if the id doesn't exist
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}
	response := models.PointsResponse{Points: curr_points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
