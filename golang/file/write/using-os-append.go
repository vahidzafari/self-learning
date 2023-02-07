package main

import (
	"fmt"
	"io"
	"os"
)

// os.OpenFile() provides a better way to create or open a file for writing. os.O_APPEND
// is saying that if the file already exists, you should append to it instead of truncating
// it. os.O_CREATE is saying that if the file does not already exist, it should be created.
// Last, os.O_WRONLY is saying that the program should open the file for writing only.

func main() {
	buffer := []byte("Data to write\n")

	f := "/tmp/append.txt"
	file, err := os.Create(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < 5; i++ {
		n, err := io.WriteString(file, string(buffer))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("wrote %d bytes\n", n)
	}

	// Append to a file
	file, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Write() needs a byte slice
	n, err := file.Write([]byte("Put some more data at the end.\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
