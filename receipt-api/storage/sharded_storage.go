package storage

import (
	"hash/fnv"
	"receipt-api/models"
)

// define number of shards
const numShards = 10

// Initialize ShardedStorage with empty maps
func NewShardedStorage() *ShardedStorage {
	storage := &ShardedStorage{}
	for i := 0; i < numShards; i++ {
		storage.shards[i] = &Shard{
			receipts: make(map[string]models.Receipt),
			points:   make(map[string]int),
		}
	}
	return storage
}

// Hash function to determine shard for a given ID
func getShardIndex(id string) int {
	hasher := fnv.New32a()                 // Fwoler-Noll-Vo 32-bit no cyrptographic hash
	hasher.Write([]byte(id))               // write the byte representation of the id
	return int(hasher.Sum32()) % numShards // computes the 32-bit hash and then modulo
}

// StoreReceipt stores a receipt and its points in the correct shard
func (s *ShardedStorage) StoreReceipt(id string, receipt models.Receipt, points int) {
	shard := s.shards[getShardIndex(id)]
	shard.mu.Lock() // write lock, one writer, blocking all reads
	defer shard.mu.Unlock()
	shard.receipts[id] = receipt
	shard.points[id] = points
}

// GetReceipt retrieves a receipt and its points from the correct shard
func (s *ShardedStorage) GetReceipt(id string) (models.Receipt, int, bool) {
	shard := s.shards[getShardIndex(id)]
	shard.mu.RLock() // read lock, multiple readers
	defer shard.mu.RUnlock()
	receipt, receiptExists := shard.receipts[id]
	points, pointsExists := shard.points[id]
	if !receiptExists || !pointsExists {
		return models.Receipt{}, 0, false
	}
	return receipt, points, true
}
