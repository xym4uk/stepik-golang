package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int
}

type Group struct {
	Students []Student
}

type Result struct {
	Average float64
}

func main() {
	var g Group
	var average float64
	//file, _ := os.Open("task.json")
	//data, _ := ioutil.ReadAll(file)
	data, _ := ioutil.ReadAll(os.Stdin)

	json.Unmarshal(data, &g)

	count := 0
	for _, student := range g.Students {
		count += len(student.Rating)
	}
	average = float64(count) / float64(len(g.Students))
	res := Result{Average: average}
	resJson, _ := json.MarshalIndent(res, "", "    ")
	fmt.Println(string(resJson))
}
