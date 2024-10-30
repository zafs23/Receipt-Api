package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"
	"receipt-api/models"
	"receipt-api/validators"
)

// ValidateReceiptMiddleware checks receipt validity before passing it to the handler
func ValidateReceiptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tempReceipt models.Receipt
		if err := json.NewDecoder(r.Body).Decode(&tempReceipt); err != nil {
			http.Error(w, "Invalid JSON request", http.StatusBadRequest)
			return
		}

		// Validate the receipt data
		if err := validators.ValidateReceipt(tempReceipt); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Re-encode the valid receipt into the request body for the handler
		bodyBytes, _ := json.Marshal(tempReceipt)
		r.Body = http.NopCloser(bytes.NewBuffer(bodyBytes))

		// Proceed to next handler
		next.ServeHTTP(w, r)
	})
}
