package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a file
	file, err := os.Create("./arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created: %v\n", file.Name())

	// Write file
	size, err := file.Write([]byte("Hello Go File"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File wrote %v bytes\n", size)

	// Read file
	fileContent, err := os.ReadFile("./arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File content: %v\n", string(fileContent))

	// Read buffer from file
	file, err = os.Open("arquivo.txt")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			fmt.Println("File readed")
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("File read: %v\n", string(buffer[:n]))
	}
	err = os.Remove("./arquivo.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
