package main

import "fmt"

func main() {
	var year int
	fmt.Scanln(&year)
	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
