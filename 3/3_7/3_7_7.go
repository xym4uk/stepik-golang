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
	//inTime = "28.07.2022 14:00:15,27.07.2022 14:00:15"
	splitted := strings.Split(inTime, ",")
	format := "02.01.2006 15:04:05"
	timeFrom, _ := time.Parse(format, splitted[0])
	timeTo, _ := time.Parse(format, splitted[1])
	if timeFrom.After(timeTo) {
		timeFrom, timeTo = timeTo, timeFrom
	}
	duration := time.Since(timeFrom).Round(time.Second) - time.Since(timeTo).Round(time.Second)
	fmt.Println(duration)
}
