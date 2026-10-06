package main

import "fmt"

func main() {
	deferTest()
}

func deferTest() {
	defer fmt.Println(1)

	fmt.Println(2)

	defer myprint(3)

	x := 4
	defer myprint(x)
	x = 5

	defer myprint(calcSix())

	fmt.Println(7)

	return
}

func myprint(i int) {
	fmt.Println(i)
}

func calcSix() int {
	fmt.Println("trying to calculate six")
	return 1 + 1 + 1 + 1 + 1 + 1
}

// Output
/*






/*
*/

/*
2
trying to calculate five
6
4
3
1
*/
