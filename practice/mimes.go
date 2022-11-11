package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var fileLink string

	fmt.Println("Enter file link: ")
	fmt.Scanf("%s", &fileLink)
	fmt.Println(fileLink)

	res, err := http.Get(fileLink)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	mimeTypes := http.DetectContentType(bytes)
	fmt.Printf("File mime type is: %s\n", mimeTypes)
}
