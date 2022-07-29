package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var inTime string
	inTime, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	inTime = strings.TrimSuffix(inTime, "\n")
	//inTime = "2020-05-15 13:00:00"
	parsed, _ := time.Parse("2006-01-02 15:04:05", inTime)
	if parsed.Hour() >= 13 {
		parsed = parsed.Add(time.Hour * 24)
	}
	fmt.Println(parsed.Format("2006-01-02 15:04:05"))
}
