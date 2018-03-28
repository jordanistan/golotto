package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD
	comb(15, 5, func(c []int) {
=======
	comb(70, 5, func(c []int) {
>>>>>>> 5566cd4dd7530de704e5e9e2581a59b05eaff927
		fmt.Println(c)
	})
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
