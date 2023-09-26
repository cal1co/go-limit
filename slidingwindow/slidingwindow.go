package slidingwindow

import (
	"errors"
	"sync"
	"time"
)

type Limiter struct {
	mu          sync.Mutex
	windowSize  time.Duration
	requests    map[time.Time]int
	maxRequests int
}

func NewSlidingWindow(windowSize time.Duration, maxRequests int) *Limiter {
	return &Limiter{
		windowSize:  windowSize,
		requests:    make(map[time.Time]int),
		maxRequests: maxRequests,
	}
}

var ErrSlidingWindowRateLimitExceeded = errors.New("rate exceeded within window")

func (limiter *Limiter) Consume() error {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	currentTime := time.Now()
	for reqTime := range limiter.requests {
		if reqTime.Add(limiter.windowSize).Before(currentTime) {
			delete(limiter.requests, reqTime)
		}
	}

	if len(limiter.requests) >= limiter.maxRequests {
		return ErrSlidingWindowRateLimitExceeded
	}

	limiter.requests[currentTime] = len(limiter.requests) + 1

	return nil
}
