package main

import (
	"fmt"
	"log"
	"net/http"
)

var PORT = ":8765"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

// mux.Handle() registers the file server as the handler for all URL paths that begin with
// /static/. However, when a match is found, we strip the /static/ prefix before the
// file server tries to serve such a request because /static/ is not part of the location
// where the actual files are located.
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", defaultHandler)

	fileServer := http.FileServer(http.Dir("/tmp/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server on:", PORT)

	err := http.ListenAndServe(PORT, mux)
	fmt.Println(err)
}
