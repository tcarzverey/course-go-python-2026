package main

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func main() {
	for {
		a = "not initialized"
		done = false
		go setup()
		for !done {
		}
		if a != "hello, world" {
			println(a)
			break
		}
	}
}
