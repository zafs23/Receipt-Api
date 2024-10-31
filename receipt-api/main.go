package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zafs23/Receipt-Api/receipt-api/handlers"
	"github.com/zafs23/Receipt-Api/receipt-api/models"
	"github.com/zafs23/Receipt-Api/receipt-api/routes"
	"github.com/zafs23/Receipt-Api/receipt-api/storage"
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
