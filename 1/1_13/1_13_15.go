package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	fmt.Scan(&n)

	res := strconv.FormatInt(n, 2)

	fmt.Println(res)
}
