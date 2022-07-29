package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int
	fmt.Scan(&a)
	slice := strconv.Itoa(a)
	fmt.Println(slice[len(slice)-1:])
}
