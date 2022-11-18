package main

import (
	"fmt"
	"net/url"
)

func main() {
	var link string

	fmt.Print("Enter an URL: ")
	fmt.Scanf("%s", &link)

	parsedUrl, err := url.Parse(link)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme: ", parsedUrl.Scheme)
	fmt.Println("Host: ", parsedUrl.Host)
	fmt.Println("Path: ", parsedUrl.Path)
}
