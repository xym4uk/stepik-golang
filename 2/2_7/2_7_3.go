package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var str []string
	fmt.Scan(&a)
	runed := []rune(a)

	for _, value := range runed {
		str = append(str, string(value))
	}

	a = strings.Join(str, "*")

	fmt.Println(a)

}
