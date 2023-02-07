package main

import (
	"bufio"
	"fmt"
	"os"
)

// UNIX uses file descriptors, which are positive integer values, as an internal
// representation for accessing open files, which is much prettier than using long paths.
// So, by default, all UNIX systems support three special and standard filenames:
// /dev/stdin, /dev/stdout, and /dev/stderr, which can also be accessed using file
// descriptors 0, 1, and 2, respectively. These three file descriptors are also called
// standard input, standard output, and standard error, respectively.

// Go uses os.Stdin for accessing standard input, os.Stdout for accessing standard
// output, and os.Stderr for accessing standard error. Although you can still use
// /dev/stdin, /dev/stdout, and /dev/stderr or the related file descriptor values
// for accessing the same devices, it is better, safer, and more portable to stick with
// os.Stdin, os.Stdout, and os.Stderr.

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	name, _ := reader.ReadString('\n')

	text := fmt.Sprintf("Given name is: %v", name)

	str := []byte(text)
	os.Stdout.Write(str)

	os.Stderr.WriteString("Error Message")
}
