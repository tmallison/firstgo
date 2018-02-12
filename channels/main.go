package main

import (
	"net/http"
	"fmt"
)

func main() {
	links := []string {
		"http://facebook.com",
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
		"http://byte.nl",
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(url string, c chan string) {
	_, error := http.Get(url)

	if error != nil {
		fmt.Println(url, "might be down I think")
		c <- url
		return
	}

	fmt.Println(url, "is up")
	c <- url
}