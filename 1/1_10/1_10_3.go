package main

import "fmt"

func main() {
	var i, j int
	sum := 0
	fmt.Scan(&i, &j)
	for i <= j {
		sum += i
		i++
	}
	fmt.Print(sum)
}
