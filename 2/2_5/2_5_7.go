package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	result := []rune(text)
	first := string(result[0:1])
	firstUpper := strings.ToUpper(first)
	length := utf8.RuneCountInString(text)
	last := string(result[length-1 : length])

	if first == firstUpper && last == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
