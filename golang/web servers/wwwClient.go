package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var timeout = time.Duration(time.Second)

// The SetDeadline() function is used by net to set the read and write deadlines of
// network connections. Due to the way the SetDeadline() function works, you need
// to call SetDeadline() before any read or write operation. Keep in mind that Go uses
// deadlines to implement timeouts, so you do not need to reset the timeout every time
// your application receives or sends any data.
func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	// The url.Parse() function parses a string into a URL structure. This means that if the
	// given argument is not a valid URL, url.Parse() is going to notice. As usual, check
	// the error variable.
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}

	t := http.Transport{
		Dial: Timeout,
	}

	c := &http.Client{
		Timeout:   15 * time.Second,
		Transport: &t,
	}

	// The http.NewRequest() function returns an http.Request object given a method, a
	// URL, and an optional body. The http.MethodGet parameter defines that we want to
	// retrieve the data using a GET HTTP method whereas URL.String() returns the string
	// value of an http.URL variable.
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	// The http.Do() function sends an HTTP request (http.Request) using an http.Client
	// and gets an http.Response. So, http.Do() does the job of http.Get() in a more
	// detailed way.
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)

	// The httputil.DumpResponse() function is used here to get the response from
	// the server and is mainly used for debugging purposes. The second argument of
	// httputil.DumpResponse() is a Boolean value that specifies whether the function is
	// going to include the body or not in its output—in our case it is set to false, which
	// excludes the response body from the output and only prints the header. If you want
	// to do the same on the server side, you should use httputil.DumpRequest().
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
