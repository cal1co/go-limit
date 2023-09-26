package tokenbucket

import (
	"errors"
	"sync"
	"time"
)

// tokens are added to a bucket at a regular interval
// buckets have a maximum capacity
// req can only be verified if there is a token
// on req accept, token is removed
// a bucket should have an id

type Bucket struct {
	id        string
	tokens    chan struct{}
	capacity  int
	rate      time.Duration
	rateMutex sync.Mutex
}

func NewTokenBucket(rate time.Duration, capacity int, id string) *Bucket {
	txns := make(chan struct{}, capacity)
	b := &Bucket{id, txns, capacity, rate, sync.Mutex{}}

	go func(b *Bucket) {
		ticker := time.NewTicker(rate)
		for range ticker.C {
			b.tokens <- struct{}{}
		}
	}(b)

	return b
}

func (b *Bucket) Fill() *Bucket {
	for i := 0; i < b.capacity; i++ {
		select {
		case b.tokens <- struct{}{}:
		default:
			return b
		}
	}
	return b
}

var ErrTokenBucketRateLimitExceeded = errors.New("bucket empty")

func (b *Bucket) Consume() error {
	select {
	case <-b.tokens:
		return nil
	default:
		return ErrTokenBucketRateLimitExceeded
	}
}

func (b *Bucket) WaitToConsume() {
	<-b.tokens
}

func (b *Bucket) GetRate() time.Duration {
	b.rateMutex.Lock()
	defer b.rateMutex.Unlock()

	rate := b.rate
	return rate
}

func (b *Bucket) BucketIsEmpty() bool {
	return len(b.tokens) == 0
}
