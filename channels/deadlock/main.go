package main

import "fmt"

func main() {
	forever := make(chan bool)
	fmt.Println("Forever")
	<-forever
}
