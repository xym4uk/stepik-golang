package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N, N)
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Print(arr[3])
}
