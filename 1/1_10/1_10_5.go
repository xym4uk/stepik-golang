package main

import "fmt"

func main() {
	var i int
	max := 0
	cnt := 0
	for fmt.Scanln(&i); i != 0; fmt.Scanln(&i) {
		if max < i {
			max = i
			cnt = 1
		} else if max == i {
			cnt++
		}
	}
	fmt.Println(cnt)
}
