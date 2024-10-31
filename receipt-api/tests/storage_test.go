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

	if points != 50 {
		t.Errorf("expected 50 points, got %d", points)
	}

	if retrievedReceipt.Retailer != receipt.Retailer {
		t.Errorf("expected retailer %s, got %s", receipt.Retailer, retrievedReceipt.Retailer)
	}
}
