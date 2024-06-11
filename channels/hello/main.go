package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go func() {
		time.Sleep(8 * time.Second)
		channel <- "Hello World"
	}()
	value := <-channel
	fmt.Println(value)
}
