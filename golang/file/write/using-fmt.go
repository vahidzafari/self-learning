package main

import (
	"fmt"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")
	file, err := os.Create("/tmp/fmt.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, string(buffer))
}
