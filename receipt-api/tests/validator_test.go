package tests

import (
	"strings"
	"testing"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/validators"
)

func TestValidateReceipt_Valid_Receipt(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.00"},
		},
		Total: "10.00",
	}

	// No error should be returned for a valid receipt
	if err := validators.ValidateReceipt(receipt); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestValidateReceipt_MissingFields(t *testing.T) {
	receipt := models.Receipt{} // Empty receipt with missing fields

	err := validators.ValidateReceipt(receipt)
	if err == nil || err.Error() != "missing required fields in the request" {
		t.Errorf("expected missing fields error, got %v", err)
	}
}

func TestValidateReceipt_InvalidDateFormat(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "01-01-2022", // Invalid date format (should be YYYY-MM-DD)
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.00"},
		},
		Total: "10.00",
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || !strings.HasPrefix(err.Error(), "invalid date format for PurchaseDate") {
		t.Errorf("expected invalid date format error, got %v", err)
	}
}

func TestValidateReceipt_InvalidTimeFormat(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "1 PM", // Invalid time format (should be HH:MM)
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.00"},
		},
		Total: "10.00",
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || !strings.HasPrefix(err.Error(), "invalid time format for PurchaseTime") {
		t.Errorf("expected invalid time format error, got %v", err)
	}
}

func TestValidateReceipt_InvalidTotalFormat(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.00"},
		},
		Total: "10.000", // Invalid total format (too many decimal places)
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || err.Error() != "invalid format for Total: value must have at most two decimal places" {
		t.Errorf("expected invalid total format error, got %v", err)
	}
}

func TestValidateReceipt_ItemInvalidPriceFormat(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Test Item", Price: "10.000"}, // Invalid price format
		},
		Total: "10.00",
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || err.Error() != "invalid format for item Price: value must have at most two decimal places" {
		t.Errorf("expected invalid item price format error, got %v", err)
	}
}

func TestValidateReceipt_EmptyShortDescription(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "", Price: "10.00"}, // Empty description
		},
		Total: "10.00",
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || err.Error() != "item ShortDescription must be between 1 and 50 characters" {
		t.Errorf("expected short description error, got %v", err)
	}
}

func TestValidateReceipt_LongShortDescription(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{
				ShortDescription: "This is a very long item description that exceeds the 1 to 50 character limit for descriptions",
				Price:            "10.00",
			},
		},
		Total: "10.00",
	}

	err := validators.ValidateReceipt(receipt)
	if err == nil || err.Error() != "item ShortDescription must be between 1 and 50 characters" {
		t.Errorf("expected long description error, got %v", err)
	}
}
