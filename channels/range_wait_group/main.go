package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go publish(channel, &waitGroup)
	go consume(channel, &waitGroup)
	waitGroup.Wait()
}

func consume(channel chan int, waitGroup *sync.WaitGroup) {
	for value := range channel {
		fmt.Printf("received value %d\n", value)
		waitGroup.Done()
	}
}

func publish(channel chan int, waitGroup *sync.WaitGroup) {
	defer close(channel)
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		fmt.Println("publish", i)
		channel <- i
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}
