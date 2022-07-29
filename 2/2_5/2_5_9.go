package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	fmt.Println(strings.Index(str1, str2))
}
