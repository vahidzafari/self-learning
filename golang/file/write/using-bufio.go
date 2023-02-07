package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buffer := []byte("Data to write\n")

	file, err := os.Create("/tmp/bufio.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	w := bufio.NewWriter(file)
	n, err := w.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}
