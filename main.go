package main

import (
	"fmt"
	"time"

	"github.com/cal1co/go-limit/tokenbucket"
)

func main() {
	fmt.Println("hello!")

	bucket := tokenbucket.NewTokenBucket(100*time.Millisecond /* refill rate*/, 10 /* capacity */, "example" /* id */)
	// bucket.Fill()

	bucket.WaitToConsume()
	fmt.Println("awaited")
}
