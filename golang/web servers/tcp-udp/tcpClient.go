package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	connect := arguments[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 	The last part of the TCP client keeps reading user input until the word STOP is given
	// as input—in this case, the client waits for the server response before terminating
	// after STOP because this is how the for loop is constructed. This mainly happens
	// because the server might have a useful answer for us, and we do not want to
	// miss that. All given user input is sent (written) to the open TCP connection using
	// fmt.Fprintf(), whereas bufio.NewReader() is used for reading data from the TCP
	// connection, just like you would do with a regular file.
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

// go run tcpConnection.go localhost:1234
