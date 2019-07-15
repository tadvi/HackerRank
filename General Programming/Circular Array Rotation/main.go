package main

import (
	"container/list"
	"fmt"
	"strings"
)

var input = strings.NewReader(`3 2 3
1 2 3
0
1
2`)

func main() {
	var n, k, q, v int
	fmt.Fscan(input, &n, &k, &q)

	res := make([]int, n)
	ls := list.New()
	for i := 0; i < n; i++ {
		fmt.Fscan(input, &v)
		ls.PushBack(v)
	}

	// Rotations.
	for i := 0; i < k; i++ {
		ls.MoveToFront(ls.Back())
	}

	// Copy to slice for faster access.
	front := ls.Front()
	res[0] = front.Value.(int)
	for i := 1; i < n; i++ {
		front = front.Next()
		res[i] = front.Value.(int)
	}

	// Output.
	var index int
	for i := 0; i < q; i++ {
		fmt.Fscan(input, &index)
		fmt.Println(res[index])
	}
}
