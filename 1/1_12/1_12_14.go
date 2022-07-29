package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	max := array[0]
	for i, value := range array {
		if value > max {
			max = array[i]
		}
	}
	fmt.Print(max)
}
