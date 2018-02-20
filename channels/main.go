package main

import (
	"net/http"
	"fmt"
	"time"
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

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down I think")
		c <- url
		return
	}

	fmt.Println(url, "is up")
	c <- url
}