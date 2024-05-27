package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Printf("Hello World, %s\n", os.Args[1])
	}
	os.Exit(0)
}
