// Problem: https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
package main

import (
	"fmt"
	"strings"
)

var input = strings.NewReader(`6 3
1 3 2 6 1 2`)

func main() {
	var n, k int
	fmt.Fscan(input, &n, &k)

	var count, v int
	bucket := make([]int, k)

	for i := 0; i < n; i++ {
		fmt.Fscan(input, &v)
		mod := v % k
		count += bucket[(k-mod)%k]
		bucket[mod]++
	}
	fmt.Println(count)
}
