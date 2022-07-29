package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x = scanner.Text()
	x = strings.Replace(x, " ", "", -1)
	x = strings.Replace(x, ",", ".", -1)
	arr := strings.Split(x, ";")
	xFloat, _ := strconv.ParseFloat(arr[0], 10)
	yFloat, _ := strconv.ParseFloat(arr[1], 10)
	fmt.Printf("%0.4f\n", xFloat/yFloat)
}
