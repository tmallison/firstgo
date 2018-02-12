package main

import (
	"net/http"
	"fmt"
	"os"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//   Same as following
	io.Copy(os.Stdout, resp.Body)
}
