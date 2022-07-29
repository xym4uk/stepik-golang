package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(fib(n))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	fibPrev, fibNext := 0, 1
	i := 1
	for fibNext <= n {
		if fibNext == n {
			return i
		}
		fibPrev, fibNext = fibNext, fibPrev+fibNext
		i++
	}
	return -1
}
