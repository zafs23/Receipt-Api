package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-api/handlers"
	"receipt-api/middleware"
	"receipt-api/models"
	"receipt-api/storage"
	"testing"

	"github.com/gorilla/mux"
)

func TestProcessReceiptHandler(t *testing.T) {
	// Initialize dependencies for ReceiptHandler
	testStorage := storage.NewShardedStorage()
	receiptCache := make(map[string]models.IDResponse)
	receiptHandler := &handlers.ReceiptHandler{
		ShardedStorage: testStorage,
		ReceiptCache:   receiptCache,
	}

	// Create a sample receipt for testing
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.00"},
		},
		Total: "10.00",
	}

	// Convert the receipt to JSON
	reqBody, _ := json.Marshal(receipt)
	req := httptest.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	ctx := context.WithValue(req.Context(), middleware.ReceiptContextKey, receipt)
	req = req.WithContext(ctx)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the ProcessReceiptHandler method
	handler := http.HandlerFunc(receiptHandler.ProcessReceiptHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Logf("Response Body: %s", rr.Body.String())
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check if response contains an ID
	var response models.IDResponse
	//json.NewDecoder(rr.Body).Decode(&response)
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if response.ID == "" {
		t.Errorf("expected a valid ID, got empty string")
	}
}

func TestGetPointsHandler(t *testing.T) {
	// Initialize dependencies for ReceiptHandler
	testStorage := storage.NewShardedStorage()
	receiptCache := make(map[string]models.IDResponse)
	receiptHandler := &handlers.ReceiptHandler{
		ShardedStorage: testStorage,
		ReceiptCache:   receiptCache,
	}

	// Assume an existing ID in storage for testing
	receiptID := "test-id"
	testReceipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        []models.Item{{ShortDescription: "Test Item", Price: "10.00"}},
		Total:        "10.00",
	}
	// Store the test receipt in sharded storage with points
	testStorage.StoreReceipt(receiptID, testReceipt, 100)

	// Create a request to get points for the test receipt ID
	req := httptest.NewRequest(http.MethodGet, "/receipts/"+receiptID+"/points", nil)
	rr := httptest.NewRecorder()

	// Setup router and register the GetPointsHandler method
	router := mux.NewRouter()
	router.HandleFunc("/receipts/{id}/points", receiptHandler.GetPointsHandler)
	router.ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}

	// Check if the response contains the correct points
	var response models.PointsResponse
	json.NewDecoder(rr.Body).Decode(&response)
	if response.Points != 100 {
		t.Errorf("expected points 100, got %d", response.Points)
	}
}
