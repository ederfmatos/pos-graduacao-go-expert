package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("Hotel booking cancelled. Timeout reached")
	case <-time.After(time.Second * 5):
		fmt.Printf("Hotel booked")
	}
}
