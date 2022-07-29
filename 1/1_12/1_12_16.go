package main

import "fmt"

func main() {
	cnt := 0
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)

	for i, _ := range arr {
		fmt.Scan(&arr[i])
		if arr[i] > 0 {
			cnt++
		}
	}
	fmt.Print(cnt)
}
