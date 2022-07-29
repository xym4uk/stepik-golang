package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; int(math.Pow(2, float64(i))) <= n; i++ {
		fmt.Printf("%.0f ", math.Pow(2, float64(i)))
	}
}
