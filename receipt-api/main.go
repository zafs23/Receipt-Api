package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-api/handlers"
	"receipt-api/models"
	"receipt-api/routes"
	"receipt-api/storage"
)

func main() {
	// Initialize sharded storage and receipt cache when the program starts
	shardedStorage := storage.NewShardedStorage()
	receiptCache := make(map[string]models.IDResponse)

	// Create an instance of ReceiptHandler with dependencies
	receiptHandler := handlers.ReceiptHandler{
		ShardedStorage: shardedStorage,
		ReceiptCache:   receiptCache,
	}

	r := routes.RegisterRoutes(&receiptHandler)
	fmt.Println("Server listening on local port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
