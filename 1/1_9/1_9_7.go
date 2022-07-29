package main

import "fmt"

func main() {
	var num string
	fmt.Scanln(&num)

	fmt.Println(string(num[0]))
}
