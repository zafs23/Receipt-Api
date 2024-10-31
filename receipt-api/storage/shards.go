package storage

import (
	"sync"

	"github.com/zafs23/Receipt-Api/receipt-api/models"
)

type Shard struct {
	mu       sync.RWMutex
	receipts map[string]models.Receipt
	points   map[string]int
}

type ShardedStorage struct {
	shards [numShards]*Shard
}
