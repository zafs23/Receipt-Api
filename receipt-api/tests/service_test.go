package tests

import (
	"receipt-api/models"
	"receipt-api/services"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        []models.Item{{ShortDescription: "Test Item", Price: "10.00"}},
		Total:        "10.00",
	}
	points, err := services.CalculatePoints(receipt)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check points according to your defined rules
	if points != 95 { // Replace with the expected points for the given receipt
		t.Errorf("expected 95 points, got %d", points)
	}
}
