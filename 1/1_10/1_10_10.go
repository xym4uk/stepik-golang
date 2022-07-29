package main

import "fmt"

func main() {
	var x, y string
	fmt.Scanln(&x, &y)

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(y); j++ {
			if x[i] == y[j] {
				fmt.Print(string(x[i]), " ")
			}
		}
	}

}
