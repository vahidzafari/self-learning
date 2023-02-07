package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

// Reading a plain text file word by word is the single most useful function that you
// want to perform on a file because you usually want to process a file on a per-word
// basis—it is illustrated in this subsection using the code found in byWord.go. The
// desired functionality is implemented in the wordByWord() function. The wordByWord()
// function uses regular expressions to separate the words found in each line of the
// input file. The regular expression defined in the regexp.MustCompile("[^\\s]+")
// statement states that we use whitespace characters to separate one word from
// another.

func wordByWord(path string) error {
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
		r := regexp.MustCompile("[^\\s]+")

		words := r.FindAllString(line, -1)

		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
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
	wordByWord(path)
}
