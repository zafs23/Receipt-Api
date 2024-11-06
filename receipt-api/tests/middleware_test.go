package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/zafs23/Receipt-Api/receipt-api/middleware"
	"github.com/zafs23/Receipt-Api/receipt-api/models"

	"testing"
)

func TestValidateReceiptMiddleware(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        []models.Item{{ShortDescription: "Test Item", Price: "10.00"}},
		Total:        "10.00",
	}

	reqBody, _ := json.Marshal(receipt)
	req := httptest.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(reqBody)) // POST request for test
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder() // response recorder

	// Dummy handler to simulate passing middleware
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	middleware.ValidateReceiptMiddleware(nextHandler).ServeHTTP(rr, req)

	// Expect the request to pass and reach the dummy handler
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}
