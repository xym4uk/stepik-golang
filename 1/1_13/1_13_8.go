package main

import "fmt"

func main() {
	var N, cnt int
	fmt.Scan(&N)
	arr := make([]int, N, N)

	for i, _ := range arr {
		fmt.Scan(&arr[i])
		if arr[i] == 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
