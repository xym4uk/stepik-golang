package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	first := num / 100
	second := num / 10 % 10
	third := num % 10

	switch {
	case first != second && second != third && first != third:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
