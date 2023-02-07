package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// The technique for reading a text file line by line is also used when reading a plain
// text file word by word as well as when reading a plain text file character by character
// because you usually process plain text files line by line. The presented utility
// prints every line that it reads, which makes it a simplified version of the cat(1) utility.

// First, you create a new reader to the desired file using a call to bufio.NewReader().
// Then, you use that reader with bufio.ReadString() in order to read the input file
// line by line. The trick is done by the parameter of bufio.ReadString(), which is a
// character that tells bufio.ReadString() to keep reading until that character is found.
// Constantly calling bufio.ReadString() when that parameter is the newline character
// (\n) results in reading the input file line by line.

func lineByLine(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one arguments!")
		return
	}
	path := arguments[1]
	lineByLine(path)
}
