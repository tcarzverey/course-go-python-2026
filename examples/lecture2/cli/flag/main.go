package main

import (
	"flag"
	"fmt"
)

func main() {
	main1()
}

func main1() {
	// simple flags
	flag.Int("n", 1, "count")

	flag.Parse()

	fmt.Printf("remaining args: %v\n", flag.Args())
	fmt.Printf("parsed flags: ")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s=%v ", f.Name, f.Value)
	})
}

func main2() {
	// multiname flags
	count := 0
	vebose := false
	flag.IntVar(&count, "count", 1, "count to print")
	flag.BoolVar(&vebose, "v", false, "verbose out")

	flag.Parse()

	fmt.Printf("remaining args: %v\n", flag.Args())
	fmt.Printf("parsed flags: ")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("%s=%v ", f.Name, f.Value)
	})

	fmt.Printf("\ncountVar: %d\n", count)
}
