package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	go func() {
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}()
	time.Sleep(time.Second * 2)
	for value := range ch {
		fmt.Printf("received value %v\n", value)
	}
}
