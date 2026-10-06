package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	fmt.Printf("args: %v\n", args)

	if len(args) < 2 {
		fmt.Println("ERROR: need at least one argument")
		os.Exit(1)
	}

	fmt.Println(strings.Join(os.Args[1:], ";"))

	return
}
