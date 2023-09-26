package main

import (
	"fmt"
	"time"

	"github.com/cal1co/go-limit/leakybucket"
)

func main() {
	fmt.Println("hello!")

	// bucket := tokenbucket.NewTokenBucket(100*time.Millisecond /* refill rate*/, 10 /* capacity */, "example" /* id */)
	// // bucket.Fill()

	// bucket.WaitToConsume()
	// fmt.Println("awaited")
	bucket := leakybucket.NewLeakyBucket(100*time.Millisecond, 5, "example")
	for i := 0; i < 20; i++ {
		// fmt.Println("bucket", bucket.IsEmpty())
		if err := bucket.Consume(); err != nil {
			fmt.Printf("Leaky Bucket: %v\n", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("consuming")
		time.Sleep(50 * time.Millisecond)
	}
}
