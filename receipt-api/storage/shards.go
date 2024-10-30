package storage

import (
	"receipt-api/models"
	"sync"
)

type Shard struct {
	mu       sync.RWMutex
	receipts map[string]models.Receipt
	points   map[string]int
}

type ShardedStorage struct {
	shards [numShards]*Shard
}
