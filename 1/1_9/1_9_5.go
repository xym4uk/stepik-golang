package main

import "fmt"

func main() {
	num := 0
	fmt.Scanln(&num)

	switch {
	case num > 0:
		fmt.Println("Число положительное")
	case num < 0:
		fmt.Println("Число отрицательное")
	case num == 0:
		fmt.Println("Ноль")
	}
}
