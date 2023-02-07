package main

import (
	"fmt"
	"io"
	"os"
)

// This subsection teaches you how to read a specific amount of data from a file. The
// presented utility can come in handy when you want to see a small part of a file. The
// numeric value that is given as a command-line argument specifies the size of the
// buffer that is going to be used for reading.

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)
	n, err := f.Read(buffer)

	// io.EOF is a special case and is treated as such
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[0:n]
}
