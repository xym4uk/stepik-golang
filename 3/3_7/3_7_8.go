package main

import (
	"fmt"
	"time"
)

func main() {
	var min, sec int
	const now int64 = 1589570165
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	start := time.Unix(now, 0)
	dur := time.Minute*time.Duration(min) + time.Second*time.Duration(sec)
	fmt.Println(start.Add(dur).UTC().Format("Mon Jan 02 15:04:05 MST 2006"))
}
