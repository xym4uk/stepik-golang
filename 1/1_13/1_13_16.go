package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	var forDelete string
	fmt.Scan(&input, &forDelete)

	fmt.Println(strings.Replace(input, forDelete, "", -1))

}
