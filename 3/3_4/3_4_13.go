package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Battery struct {
	charge string
}

func (b Battery) String() string {
	str := strings.Replace(b.charge, "0", "", -1)
	len := utf8.RuneCountInString(str)
	runed := make([]rune, 0, 10)
	for i := 0; i < 10; i++ {
		if i < 10-len {
			runed = append(runed, rune(32))
		} else {
			runed = append(runed, rune(88))
		}
	}

	return fmt.Sprintf("[%v]", string(runed))
}

func main() {
	var charge string
	fmt.Scan(&charge)
	batteryForTest := Battery{charge: charge}