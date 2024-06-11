package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	for workerId := range 10_000 {
		go worker(workerId, channel)
	}
	for i := 0; i < 100_000; i++ {
		channel <- i
	}
}

func worker(id int, channel chan int) {
	for value := range channel {
		fmt.Printf("worker %d received %d\n", id, value)
		time.Sleep(time.Second)
	}
}
