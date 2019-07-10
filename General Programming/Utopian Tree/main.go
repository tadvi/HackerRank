// Problem:
package main

import (
	"fmt"
	"strings"
)

var input = strings.NewReader(`3
0
1
4`)

func main() {
	var n, max int
	fmt.Fscan(input, &n)

	t := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &t[i])
		if t[i] > max {
			max = t[i]
		}
	}

	cache := make([]int, max+1)
	cache[0] = 1

	for i := 1; i <= max; i++ {
		if i%2 == 1 {
			cache[i] = cache[i-1] * 2
		} else {
			cache[i] = cache[i-1] + 1
		}
	}

	for _, v := range t {
		fmt.Println(cache[v])
	}
}
