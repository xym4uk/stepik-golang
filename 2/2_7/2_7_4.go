package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	splitted := strings.Split(a, "")
	max := splitted[0]
	for _, val := range splitted {
		if max < val {
			max = val
		}
	}
	fmt.Println(max)

}
