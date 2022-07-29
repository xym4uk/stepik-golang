package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	//1 градус = 6 минут
	hours := n / 30
	minutes := 2 * (n % 30)
	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
