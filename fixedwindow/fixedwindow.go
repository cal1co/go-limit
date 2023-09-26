package fixedwindow

import (
	"errors"
	"sync"
	"time"
)

type Limiter struct {
	mu           sync.Mutex
	windowStart  time.Time
	requestCount int
	windowSize   time.Duration
	maxRequests  int
}

func NewFixedWindow(windowSize time.Duration, maxRequests int) *Limiter {
	return &Limiter{
		windowStart:  time.Now(),
		requestCount: 0,
		windowSize:   windowSize,
		maxRequests:  maxRequests,
	}
}

var ErrFixedWindowRateLimitExceeded = errors.New("rate limit exceeded")

func (limiter *Limiter) Consume() error {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	currentTime := time.Now()

	if currentTime.Sub(limiter.windowStart) >= limiter.windowSize {
		limiter.windowStart = currentTime
		limiter.requestCount = 0
	}

	if limiter.requestCount >= limiter.maxRequests {
		return ErrFixedWindowRateLimitExceeded
	}

	limiter.requestCount++

	return nil
}
