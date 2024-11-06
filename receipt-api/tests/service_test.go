package tests

import (
	"testing"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/services"
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

// test that GenerateReceiptID produces a unique hash ID for a given receipt
func TestGenerateReceiptID(t *testing.T) {
	// Define a sample receipt
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        []models.Item{{ShortDescription: "Test Item", Price: "10.00"}},
		Total:        "10.00",
	}

	// Generate ID for the first time
	id1 := services.GenerateReceiptID(receipt)

	// Generate ID for the second time with the same data
	id2 := services.GenerateReceiptID(receipt)

	// Assert that the ID generated for the same receipt data is consistent
	if id1 != id2 {
		t.Errorf("expected ID to be consistent ID, got %s and %s", id1, id2)
	}

	// Define a different receipt to test uniqueness of the generated ID
	receiptDifferent := models.Receipt{
		Retailer:     "Another Retailer",
		PurchaseDate: "2023-01-01",
		PurchaseTime: "14:01",
		Items:        []models.Item{{ShortDescription: "Different Item", Price: "15.00"}},
		Total:        "15.00",
	}

	// Generate ID for different data
	id3 := services.GenerateReceiptID(receiptDifferent)

	// Assert that the ID for different receipt data is unique
	if id1 == id3 {
		t.Errorf("expected different IDs for different receipt data, got identical IDs %s and %s", id1, id3)
	}
}
