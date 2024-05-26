package main

import "fmt"

func main() {
	fmt.Println("Hello 1")
	defer fmt.Println("Hello 2")
	fmt.Println("Hello 3")
	defer fmt.Println("Hello 4")
	fmt.Println("Hello 5")
	defer fmt.Println("Hello 6")
}
