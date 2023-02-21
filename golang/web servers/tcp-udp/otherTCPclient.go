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
		fmt.Println("Please provide a server:port string!")
		return
	}

	// The net.ResolveTCPAddr() function is specific to TCP connections, hence its name,
	// and resolves the given address to a *net.TCPAddr value, which is a structure that
	// represents the address of a TCP endpoint—in this case, the endpoint is the TCP
	// server we want to connect to.
	connect := arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		fmt.Println("ResolveTCPAddr:", err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
	}
}
