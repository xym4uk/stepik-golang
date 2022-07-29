package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	first := num / 100
	second := num / 10 % 10
	last := num % 10
	fmt.Printf("%d%d%d", last, second, first)
}
