package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	for i, v := range str1 {
		if i%2 != 0 {
			fmt.Printf("%c", v)
		}
	}
}
