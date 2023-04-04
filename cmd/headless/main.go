package main

import (
	"fmt"

	"gopkg.in/headzoo/surf.v1"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	// Outputs: "The Go Programming Language"
	fmt.Println(bow.Title())
}
