package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runed := []rune(text)

	flag := true
	for idx, _ := range runed {
		if runed[idx] != runed[len(runed)-1-idx] {
			flag = false
		}
	}
	if flag {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
