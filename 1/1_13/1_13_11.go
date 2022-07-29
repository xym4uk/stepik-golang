package main

import "fmt"

func main() {
	var a, b, max int
	fmt.Scan(&a, &b)

	max = 1
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			max = i
			break
		}
	}
	if max == 1 {
		fmt.Print("NO")
	} else {
		fmt.Print(max)
	}
}
