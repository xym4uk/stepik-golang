package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	first := num / 1000
	last := num % 1000

	sumFirst := first/100 + first/10%10 + first%10
	sumLast := last/100 + last/10%10 + last%10

	if sumFirst == sumLast {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
