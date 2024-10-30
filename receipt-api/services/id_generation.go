package services

import "github.com/google/uuid"

func GenerateReceiptID() string {
	return uuid.New().String()
}
