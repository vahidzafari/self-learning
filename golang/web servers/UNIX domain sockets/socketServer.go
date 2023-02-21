package main

import (
	"fmt"
	"net"
	"os"
)

// In this second part of echo(), we send back to the client the data that the client sent
// to us. The buf[0:n] notation makes sure that we are going to send back the same
// amount of data that was read even if the size of the buffer is bigger.
// This function serves all client connections—as you are going to see in a while, it is
// executed as a goroutine, which is the main reason that it does not return any values.
func echo(c net.Conn) {
	for {
		buf := make([]byte, 128)

		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Read:", err)
			return
		}

		data := buf[0:n]
		fmt.Print("Server got: ", string(data))

		_, err = c.Write(data)
		if err != nil {
			fmt.Println("Write:", err)
			return
		}
	}
}

// You cannot tell whether this function serves TCP/IP connections or UNIX socket
// domain connections, which mainly happens because UNIX treats all connections as
// files.
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need socket path")
		return
	}
	socketPath := os.Args[1]

	// If the socket file already exists, you should delete it before the program continues—
	// net.Listen() creates that file again.
	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing", socketPath)
		err := os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// What makes this a UNIX domain socket server is the use of net.Listen() with the
	// "unix" parameter. In this case, we need to provide net.Listen() with the path of the
	// socket file.
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		go echo(fd)
	}
}
