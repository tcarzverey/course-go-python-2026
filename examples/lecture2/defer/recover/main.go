package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	recoverAndContinue(panickingFunc)
}

func recoverAndContinue(f func()) {
	tryFunc := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: \"%v\"\n", r)
			}
		}()

		f()
	}

	for {
		tryFunc()
	}
}

func panickingFunc() {
	if rand.Intn(2) == 0 {
		panic("oh no!")
	} else {
		fmt.Println("all good, sleeping!")
		time.Sleep(time.Second * 2)
	}
}
