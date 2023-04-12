package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: no file argument")
		os.Exit(1)
	}

	file, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Println("Error: reading file", os.Args[1])
		fmt.Println(error)
		os.Exit(2)
	}

	io.Copy(os.Stdout, file)

}
