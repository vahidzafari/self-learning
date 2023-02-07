package main

import (
	"fmt"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")
	file, err := os.Create("/tmp/os.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer file.Close()
	n, err := file.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\n", n)
}
