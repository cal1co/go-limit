package leakybucket

import (
	"errors"
	"sync"
	"time"
)

type Bucket struct {
	id        string
	tokens    chan struct{}
	capacity  int
	rate      time.Duration
	rateMutex sync.Mutex
}

func NewLeakyBucket(rate time.Duration, capacity int, id string) *Bucket {
	txns := make(chan struct{}, capacity)
	b := &Bucket{id, txns, capacity, rate, sync.Mutex{}}

	go func(b *Bucket) {
		ticker := time.NewTicker(rate)
		for range ticker.C {
			<-b.tokens
		}
	}(b)

	return b
}

var ErrLeakyBucketRateLimitExceeded = errors.New("bucket full")

func (b *Bucket) Consume() error {
	select {
	case b.tokens <- struct{}{}:
		return nil
	default:
		return ErrLeakyBucketRateLimitExceeded
	}
}

func (b *Bucket) GetRate() time.Duration {
	b.rateMutex.Lock()
	defer b.rateMutex.Unlock()

	rate := b.rate
	return rate
}

func (b *Bucket) SetRate(rate time.Duration) {
	b.rateMutex.Lock()
	defer b.rateMutex.Unlock()

	b.rate = rate
}

func (b *Bucket) IsEmpty() bool {
	return len(b.tokens) == 0
}

func (b *Bucket) IsFull() bool {
	return len(b.tokens) == b.capacity
}
