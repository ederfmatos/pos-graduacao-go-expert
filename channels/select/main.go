package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	for range 2 {
		select {
		case msg := <-c1:
			println(msg)
		case msg := <-c2:
			println(msg)
		case <-time.After(time.Second * 2):
			println("timeout")
		}
	}

}
