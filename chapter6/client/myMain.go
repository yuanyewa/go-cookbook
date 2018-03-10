package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("this is from go")
	c := http.DefaultClient
	resp0, _ := http.Get("https://www.golang.org")
	resp1, _ := c.Get("http://www.google.com")
	fmt.Println(resp0.StatusCode)
	fmt.Println(resp1.StatusCode)
	fmt.Println(resp1.ContentLength)
}