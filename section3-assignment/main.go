package main

import (
	"fmt"
)

func main() {
	s := intSlcZeroTo(10)

	for _, n := range s {
		msg := fmt.Sprintf("%v is %v", n, oddOrEven(n))
		fmt.Println(msg)
	}
}

func oddOrEven(i int) string {
	if i%2 == 0 {
		return "Even"
	}

	return "Odd"
}

func intSlcZeroTo(n int) []int {
	s := []int{}

	for i := 0; i <= n; i++ {
		s = append(s, i)
	}

	return s
}
