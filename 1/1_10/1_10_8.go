package main

import "fmt"

func main() {
	var num int
	for {
		fmt.Scanln(&num)
		if num < 10 {
			continue
		}
		if num > 100 {
			break
		}
		fmt.Println(num)
	}
}
