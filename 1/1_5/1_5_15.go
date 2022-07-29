package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	ten := n % 100
	digit := ten / 10
	fmt.Println(digit)
}
