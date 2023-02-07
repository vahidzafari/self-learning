package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// In this subsection, you learn how to read a text file character by character, which is a
// pretty rare requirement unless you want to develop a text editor. You take each line
// that you read and split it using a for loop with range, which returns two values. You
// discard the first, which is the location of the current character in the line variable,
// and you use the second. However, that value is a rune, which means that you have
// to convert it into a character using string().

func charByChar(path string) error {
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
			return err
		}
		for _, x := range line {
			fmt.Println(string(x))
		}
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
	charByChar(path)
}
