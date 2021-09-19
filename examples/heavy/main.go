package main

import (
	"fmt"
	"os"

	"github.com/sarir/go-link"
)

func main() {
	f, err := os.Open("go-awesome.html")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	links, err := link.Parse(f)

	if err != nil {
		panic(err)
	}

	for i, l := range links {
		fmt.Printf("%d) (%s) %s\n", i, l.Href, l.Text)
	}
}
