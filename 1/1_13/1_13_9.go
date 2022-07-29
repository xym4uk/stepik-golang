package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	arr := make([]int, N, N)

	cnt := 0
	min := 10000000
	for i, _ := range arr {
		fmt.Scan(&arr[i])
		if arr[i] < min {
			min = arr[i]
			cnt = 1
		} else if arr[i] == min {
			cnt++
		}
	}
	fmt.Print(cnt)
}
