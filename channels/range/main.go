package main

import "fmt"

func main() {
	channel := make(chan int)
	go publish(channel)
	consume(channel)
}

func consume(channel chan int) {
	for value := range channel {
		fmt.Printf("received value %d\n", value)
	}
}

func publish(channel chan int) {
	defer close(channel)
	for i := 0; i < 10; i++ {
		fmt.Println("publish", i)
		channel <- i
	}
}
