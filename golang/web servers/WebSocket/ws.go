package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// The websocket.Upgrader method of the gorilla/websocket package upgrades
// an HTTP server connection to the WebSocket protocol and allows you to define
// the parameters of the upgrade. After that, your HTTP connection is a WebSocket
// connection, which means that you will not be allowed to execute statements that
// work with the HTTP protocol.

// Last, remember that in a WebSocket connection, you cannot use fmt.Fprintf()
// statements for sending data to the WebSocket client—if you use any of these, or
// any other call that can implement the same functionality, the WebSocket connection
// fails and you are not going to be able to send or receive any data. Therefore, the
// only way to send and receive data in a WebSocket connection implemented with
// gorilla/websocket is through WriteMessage() and ReadMessage() calls, respectively.

var PORT = ":1234"
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket!")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Connection from:", r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader.Upgrade:", err)
		return
	}
	defer ws.Close()

	// 	The for loop in wsHandler() handles all incoming messages for /ws—you can use
	// any technique you want. Additionally, in the presented implementation, only the
	// client is allowed to close an existing WebSocket connection unless there is a network
	// issue, or the server process is killed.
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("From", r.Host, "read", err)
			break
		}
		log.Print("Received: ", string(message))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("WriteMessage:", err)
			break
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/ws", http.HandlerFunc(wsHandler))

	log.Println("Listening to TCP Port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
