package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// NOTE: this is contrived!
	// manually make a slice with 99,999 spaces in it
	// bs := make([]byte, 99999)
	// Read will not add more space to a slice if it fills up
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// this accomplishes the same thing as above, but with the Writer interface
	// take a byte slice and write it out somewhere, in this case Stdout
	// io.Copy first arg is something that implements Writer
	// io.Copy second arg is something that implements Reader
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// our custom logWriter type now implements the Writer interface
// because it implements a function Write that accepts a byte slice and returns and (int, error)
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
