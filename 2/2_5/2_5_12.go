package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scan(&str)
	runed := []rune(str)
	flag := true
	if len(runed) < 5 {
		flag = false
	}
	for _, val := range runed {
		if !(unicode.Is(unicode.Latin, val) || unicode.IsNumber(val)) {
			flag = false
		}
	}
	if flag {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
