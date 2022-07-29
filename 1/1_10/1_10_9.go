package main

import "fmt"

func main() {
	var x, p, y, years int
	fmt.Scanln(&x)
	fmt.Scanln(&p)
	fmt.Scanln(&y)
	for x < y {
		years++
		x += x * p / 100
	}
	fmt.Println(years)
}
