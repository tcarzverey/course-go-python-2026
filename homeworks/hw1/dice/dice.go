package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const (
	minValue = 1
	maxValue = 100
)

// diceSpec описывает бросок N костей с M гранями.
type diceSpec struct {
	count int
	sides int
}

// parseSpec разбирает аргумент вида NdM.
func parseSpec(s string) (diceSpec, error) {
	ns, ms, ok := strings.Cut(s, "d")
	if !ok {
		return diceSpec{}, fmt.Errorf("invalid dice format %q, expected NdM", s)
	}

	n, err := parseValue(s, ns, "count", minValue, maxValue)
	if err != nil {
		return diceSpec{}, err
	}

	m, err := parseValue(s, ms, "sides", minValue, maxValue)
	if err != nil {
		return diceSpec{}, err
	}

	return diceSpec{count: n, sides: m}, nil
}

// parseValue разбирает число из диапазона [lo, hi].
func parseValue(spec, value, name string, lo, hi int) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid dice format %q: %w", spec, err)
	}
	if v < lo || v > hi {
		return 0, fmt.Errorf("dice %s %d out of range [%d, %d]", name, v, lo, hi)
	}
	return v, nil
}

// roll возвращает броски костей: каждое значение от 1 до числа граней.
func roll(r *rand.Rand, spec diceSpec) []int {
	values := make([]int, spec.count)
	for i := range values {
		values[i] = r.Intn(spec.sides) + 1
	}

	return values
}

// sumValues возвращает сумму значений костей.
func sumValues(values []int) int {
	s := 0
	for _, v := range values {
		s += v
	}

	return s
}
