package main

import "fmt"

func main() {
	var n int
	var str string
	fmt.Scan(&n)

	switch {
	case (n < 10 || n > 20) && (n%10 == 2 || n%10 == 3 || n%10 == 4):
		str = "korovy"
	case n == 11:
		str = "korov"
	case n%10 == 1:
		str = "korova"
	default:
		str = "korov"
	}

	fmt.Printf("%d %s", n, str)
}
