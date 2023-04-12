package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	// fmt.Printf("response: %v\n", response) // headers, status code etc, but not body
	// io.Copy(os.Stdout, response.Body)

	lw := logWriter{}
	io.Copy(lw, response.Body)

	// byteSlice := make([]byte, 99999) // guess how much data can there be
	// response.Body.Read(byteSlice)    // reader outputs the read data to provided byte slice
	// fmt.Println(byteSlice)
	// fmt.Println(string(byteSlice))

}

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote", len(bs), "bytes")
	return len(bs), nil
}
