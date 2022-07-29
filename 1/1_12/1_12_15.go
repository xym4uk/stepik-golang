package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i, _ := range arr {
		fmt.Scan(&arr[i])
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}
