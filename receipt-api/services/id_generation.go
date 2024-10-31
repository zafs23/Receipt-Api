package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
)

func GenerateReceiptID(receipt models.Receipt) string {
	// Marshal receipt data to JSON
	data, _ := json.Marshal(receipt)

	// Create SHA-256 hash of the JSON data
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
	//return uuid.New().String()
}
