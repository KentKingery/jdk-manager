package main

import (
	
	"flag"
	"fmt"
	// "io"
	"os"
)

func main() {
	fmt.Println(flag.Args())
	showPath := flag.Bool("path", false, "Show message")
	// Parsing the flags provided by the user
	flag.Parse()
	if (*showPath) {
		fmt.Println(os.Getenv("PATH"))
	}
	
}