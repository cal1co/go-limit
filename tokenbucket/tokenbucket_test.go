package tokenbucket

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBucketIsEmpty(t *testing.T) {
	rate := 100 * time.Millisecond
	capacity := 5
	id := "testBucket"
	bucket := NewTokenBucket(rate, capacity, id)

	isEmpty := bucket.IsEmpty()
	assert.True(t, isEmpty, "Bucket should be empty initially.")
}

func TestGetRate(t *testing.T) {
	rate := 100 * time.Millisecond
	capacity := 5
	id := "testBucket"
	bucket := NewTokenBucket(rate, capacity, id)

	rateValue := bucket.GetRate()
	assert.Equal(t, rate, rateValue, "GetRate() should return the expected rate.")
}
