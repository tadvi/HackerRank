// Problem: https://www.hackerrank.com/challenges/plus-minus/problem
package main

import (
	"fmt"
	"strings"
)

var input = strings.NewReader(`6
-4 3 -9 0 4 1 `)

func main() {
	var size, n int
	var pos, zero, neg float64
	fmt.Fscan(input, &size)

	for {
		_, err := fmt.Fscan(input, &n)
		if err != nil {
			break
		}

		switch {
		case n < 0:
			neg++
		case n == 0:
			zero++
		case n > 0:
			pos++
		}
	}

	fmt.Printf("%.6f\n", pos/float64(size))
	fmt.Printf("%.6f\n", neg/float64(size))
	fmt.Printf("%.6f\n", zero/float64(size))
}
