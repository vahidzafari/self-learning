package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// The TCP server, which uses net.Listen(), returns the current date and time to the
// client in a single network packet. In practice, this means that after accepting a
// client connection, the server gets the time and date from the operating system and
// sends that data back to the client. The net.Listen() function listens for
// connections, whereas the net.Accept() method waits for the next connection and
// returns a generic Conn variable with the client information.

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	// The net.Listen() function listens for connections and is what makes that particular
	// program a server process. If the second parameter of net.Listen() contains a port
	// number without an IP address or a hostname, net.Listen() listens to all available
	// IP addresses of the local system, which is the case here.
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	// We just call Accept() and wait for a client connection—Accept() blocks until a
	// connection comes. There is something unusual with this particular TCP server: it can
	// only serve the first TCP client that is going to connect to it because the Accept() call
	// is outside of the for loop and therefore is called only once. Each individual client
	// should be specified by a different Accept() call.
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}
		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
