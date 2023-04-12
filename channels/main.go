package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	// for { // infinite loop (but always awaits for a channel to receive by <-c)
	// 	go checkUrl(<-c, c)
	// }

	// equivalent but more clear
	// for url := range c {
	// 	go checkUrl(url, c)
	// }

	// function literal with sleeping between next call
	for url := range c {
		go func(copiedUrl string) {
			time.Sleep(2 * time.Second) // 2s sleep
			checkUrl(copiedUrl, c)
		}(url)

		// nested function has access to copy of the object instead of the same instance
	}
}

func checkUrl(url string, c chan string) {
	// time.Sleep(time.Second)
	response, error := http.Get(url)

	if error != nil || response.StatusCode != 200 {
		fmt.Println(url, "is down")
		c <- url
		return
	}

	fmt.Println(url, "is up")
	c <- url
}
