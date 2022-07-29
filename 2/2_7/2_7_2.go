package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(calc(a, b))
}

func calc(a, b float64) float64 {
	return math.Sqrt((a * a) + (b * b))
}
