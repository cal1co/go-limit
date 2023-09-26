package main

import (
	"fmt"
	"time"

	"github.com/cal1co/go-limit/tokenbucket"
)

func main() {

	bucket := tokenbucket.NewTokenBucket(100*time.Millisecond, 5, "example")
	for i := 0; i < 15; i++ {
		if err := bucket.ConsumeTokens(1); err != nil {
			fmt.Printf("error: %v\n", err)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
