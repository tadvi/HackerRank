// Problem: https://www.hackerrank.com/challenges/staircase/problem
package main

import (
	"fmt"
	"strings"
)

var input = strings.NewReader(`6`)

func main() {
	var n int
	fmt.Fscan(input, &n)

	for i := 1; i < n+1; i++ {
		fmt.Print(strings.Repeat(" ", n-i))
		fmt.Print(strings.Repeat("#", i), "\n")
	}
}
