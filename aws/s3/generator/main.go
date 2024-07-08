package main

import (
	"fmt"
	"os"
)

func main() {
	for i := range 9999 {
		file, err := os.Create(fmt.Sprintf("aws/s3/tmp/%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString(fmt.Sprintf("Hello World: %d\n", i))
	}
}
