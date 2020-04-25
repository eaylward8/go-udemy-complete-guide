package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args returns a string slice with cmd line arguments
	f, err := openFile(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	printFile(f)
}

func openFile(p string) (*os.File, error) {
	return os.Open(p)
}

func printFile(f *os.File) {
	fmt.Println(io.Copy(os.Stdout, f))
}
