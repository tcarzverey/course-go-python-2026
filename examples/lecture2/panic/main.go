package main

import "os"

func main() {
	internal()

	_, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
}

func internal() {
	panic("a problem")
}
