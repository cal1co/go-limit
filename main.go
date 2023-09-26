package main

import (
	"fmt"
	"time"

	"github.com/cal1co/go-limit/tokenbucket"
)

func main() {
	fmt.Println("hello!")

	bucket := tokenbucket.NewTokenBucket(100*time.Millisecond, 10, "example")
	// bucket.Fill()

	bucket.WaitToConsume()
	fmt.Println("awaited")
}
