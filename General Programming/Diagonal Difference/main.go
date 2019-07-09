// Problem: https://www.hackerrank.com/challenges/diagonal-difference/problem
package main

import (
	"fmt"
	"math"
	"strings"
)

var input = strings.NewReader(`3
11 2 4
4 5 6
10 8 -12 `)

func main() {
	var n int
	fmt.Fscan(input, &n)

	var v, pdiag, sdiag int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(input, &v)
			if i == j {
				pdiag += v
			}
			if i+j == n-1 {
				sdiag += v
			}
		}
	}

	diff := int(math.Abs(float64(pdiag - sdiag)))
	fmt.Println(diff)
}
