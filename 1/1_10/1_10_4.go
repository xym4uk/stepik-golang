package main

import "fmt"

func main() {
	var count, number int
	sum := 0
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		fmt.Scan(&number)
		if number%8 == 0 && number > 8 && number < 100 {
			sum += number
		}
	}
	fmt.Print(sum)
}
