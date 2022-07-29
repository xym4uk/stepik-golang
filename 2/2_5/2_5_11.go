package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	strPtr := &str
	for _, value := range str {
		if strings.Count(str, string(value)) > 1 {
			*strPtr = strings.Replace(str, string(value), "", -1)
		}
	}
	fmt.Println(*strPtr)
}
