package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// The filepath.Base() function returns the last element of a path. When given
	// os.Args[0] as its parameter, it returns the name of the executable binary file.
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	// W e get the URL and get its data using http.Get(), which returns an *http.Response
	// and an error variable. The *http.Response value contains all the information so
	// you do not need to make any additional calls to http.Get().
	URL := os.Args[1]
	data, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The io.Copy() function reads from the data.Body reader, which contains the body of
	// the server response, and writes the data to os.Stdout. As os.Stdout is always open,
	// you do not need to open it for writing. Therefore, all data is written to standard
	// output, which is usually the terminal window.
	_, err = io.Copy(os.Stdout, data.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	data.Body.Close()
}
