package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/validators"
)

// define a custom type for context keys
type contextKey string

const ReceiptContextKey contextKey = "validatedReceipt" // define a key for storing receipt in context

func ValidateReceiptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validateReceipt(w, r, next)
	})
}

// validateReceipt performs validation and proceeds to next handler if valid
func validateReceipt(w http.ResponseWriter, r *http.Request, next http.Handler) {
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

	// Store the validated receipt in the request context
	ctx := context.WithValue(r.Context(), ReceiptContextKey, tempReceipt)

	// Pass the modified request with context to the next handler
	next.ServeHTTP(w, r.WithContext(ctx))
}
