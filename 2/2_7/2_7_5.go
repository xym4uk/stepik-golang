package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, res string
	fmt.Scan(&a)
	splitted := strings.Split(a, "")
	for _, val := range splitted {
		num, _ := strconv.ParseInt(val, 10, 10)
		res = res + strconv.FormatInt(num*num, 10)
	}
	fmt.Println(res)

}
