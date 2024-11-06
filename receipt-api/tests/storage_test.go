package tests

import (
	"testing"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/storage"
)

func TestStoreAndRetrieveReceipt(t *testing.T) {
	shardedStorage := storage.NewShardedStorage()
	receiptID := "test-id"
	receipt := models.Receipt{
		Retailer:     "Test Retailer",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items:        []models.Item{{ShortDescription: "Test Item", Price: "10.00"}},
		Total:        "10.00",
	}
	shardedStorage.StoreReceipt(receiptID, receipt, 50)

	// Retrieve
	retrievedReceipt, points, found := shardedStorage.GetReceipt(receiptID)
	if !found {
		t.Fatal("receipt not found in storage")
	}

	// check if all fields match
	if points != 50 {
		t.Errorf("expected 50 points, got %d", points)
	}

	if retrievedReceipt.Retailer != receipt.Retailer {
		t.Errorf("expected retailer %s, got %s", receipt.Retailer, retrievedReceipt.Retailer)
	}

	if retrievedReceipt.PurchaseDate != receipt.PurchaseDate {
		t.Errorf("expected purchase date %s, got %s", receipt.PurchaseDate, retrievedReceipt.PurchaseDate)
	}

	if retrievedReceipt.PurchaseTime != receipt.PurchaseTime {
		t.Errorf("expected purchase time %s, got %s", receipt.PurchaseTime, retrievedReceipt.PurchaseTime)
	}

	if len(retrievedReceipt.Items) != len(receipt.Items) {
		t.Fatalf("expected %d items, got %d", len(receipt.Items), len(retrievedReceipt.Items))
	}
	for i, item := range retrievedReceipt.Items {
		if item.ShortDescription != receipt.Items[i].ShortDescription {
			t.Errorf("expected item description %s, got %s", receipt.Items[i].ShortDescription, item.ShortDescription)
		}
		if item.Price != receipt.Items[i].Price {
			t.Errorf("expected item price %s, got %s", receipt.Items[i].Price, item.Price)
		}
	}

}
