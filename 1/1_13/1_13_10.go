package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Print(1 + (num-1)%9)
}
