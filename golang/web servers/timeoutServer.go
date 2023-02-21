package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

// This section presents a technique for timing out network connections that take too
// long to finish on the server side. This is much more important than the client side
// because a server with too many open connections might not be able to process more
// requests unless some of the already open connections close. This usually happens
// for two reasons. The first reason is software bugs, and the second reason is when a
// server is experiencing a Denial of Service (DoS) attack!
func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)
	m := http.NewServeMux()

	// This is where the timeout periods are defined. Note that you can define timeout
	// periods for both reading and writing processes. The value of the ReadTimeout field
	// specifies the maximum duration allowed to read the entire client request, including
	// the body, whereas the value of the WriteTimeout field specifies the maximum time
	// duration before timing out the sending of the client response.
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
