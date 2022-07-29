package main

import (
	"fmt"
	"time"
)

func main() {
	var inTime string
	fmt.Scan(&inTime)
	parsed, _ := time.Parse("2006-01-02T15:04:05-07:00", inTime)
	fmt.Println(parsed.Format(time.UnixDate))
}
