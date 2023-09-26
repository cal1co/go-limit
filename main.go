package main

import (
	"fmt"
	"time"

	"github.com/cal1co/go-limit/fixedwindow"
)

func main() {
	fmt.Println("hello!")

	// bucket := tokenbucket.NewTokenBucket(100*time.Millisecond /* refill rate*/, 10 /* capacity */, "example" /* id */)
	// // bucket.Fill()

	// bucket.WaitToConsume()
	// fmt.Println("awaited")
	limiter := fixedwindow.NewFixedWindow(5*time.Second, 3)

	for i := 0; i < 10; i++ {
		if err := limiter.Consume(); err != nil {
			fmt.Printf("fixed window: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Println("consume success")

		time.Sleep(1 * time.Second)
	}
}
