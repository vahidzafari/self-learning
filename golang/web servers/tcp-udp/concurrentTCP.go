package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// This section teaches a pattern for developing concurrent TCP servers, which are
// servers that are using separate goroutines to serve their clients following a successful
// Accept() call. Therefore, such servers can serve multiple TCP clients at the same
// time. This is how real-world production servers and services are implemented.

var count = 0

func handleConnection(c net.Conn) {
	fmt.Print(".")

	// 	The for loop makes sure that handleConnection() is not going to exit automatically.
	// Once again, the STOP keyword stops the current client connection—however, the
	// server process, as well as all other active client connections, are going to keep
	// running.
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)

		counter := "Client number: " + strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
