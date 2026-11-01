package main

import (
	"fmt"
	"strings"
)

// formatRolls формирует строку Rolls.
func formatRolls(values []int) string {
	var b strings.Builder
	b.WriteString("Rolls: ")

	for i, v := range values {
		if i > 0 {
			b.WriteByte('+')
		}
		fmt.Fprintf(&b, "%d", v)
	}

	if len(values) > 1 {
		fmt.Fprintf(&b, "=%d", sumValues(values))
	}

	return b.String()
}

// printDetailed выводит подробный разбор итерации: кости, сумму и среднее.
func printDetailed(pad string, values []int) {
	for i, v := range values {
		fmt.Printf("%sDice %d: %d\n", pad, i+1, v)
	}

	total := sumValues(values)
	fmt.Printf("%sSum: %d\n", pad, total)
	fmt.Printf("%sAverage: %g\n", pad, float64(total)/float64(len(values)))
}
