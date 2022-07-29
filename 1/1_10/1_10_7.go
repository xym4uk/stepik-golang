package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scanln(&n)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	for i := 1; i <= n; i++ {
		if i%d != 0 && i%c == 0 {
			fmt.Println(i)
			break
		}
	}
}
