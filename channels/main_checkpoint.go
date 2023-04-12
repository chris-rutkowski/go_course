package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// for _, url := range urls {
	// 	// there is the main routine
	// 	// go keyword invokes a function on a separate routine AKA thread
	// 	go checkUrlWithoutChannel(url)
	// 	// you may not see the Println inside checkUrl
	// 	// because main routine exits before responses has a chance to arrive
	// }

	c := make(chan string) // channel to communicate between routines

	for _, url := range urls {
		go checkUrl(url, c)
	}

	// // first message and program exists
	// fmt.Println(<-c)

	// // so adding remaining four awaiters
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)

	// // but this will wait forever
	// fmt.Println(<-c)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

// func checkUrlWithoutChannel(url string) {
// 	response, error := http.Get(url)

// 	if error != nil || response.StatusCode != 200 {
// 		fmt.Println(url, "is down")
// 		return
// 	}

// 	fmt.Println(url, "is up")
// }

func checkUrl(url string, c chan string) {
	response, error := http.Get(url)

	if error != nil || response.StatusCode != 200 {
		c <- url + " is down (c)"
		return
	}

	c <- url + "is up (c)"
}
