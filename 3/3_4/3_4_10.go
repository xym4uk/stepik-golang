package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask()

	switch val := value1.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T\n", value1, val)
		return
	}

	switch val := value2.(type) {
	case float64:
		break
	default:
		fmt.Printf("value=%v: %T\n", value2, val)
		return
	}

	switch operation.(type) {
	case string:
		switch operation {
		case "+":
			fmt.Printf("%.4f", value1.(float64)+value2.(float64))
		case "-":
			fmt.Printf("%.4f", value1.(float64)-value2.(float64))
		case "/":
			fmt.Printf("%.4f", value1.(float64)/value2.(float64))
		case "*":
			fmt.Printf("%.4f", value1.(float64)*value2.(float64))
		}
	default:
		fmt.Println("неизвестная операция")
		return
	}

}
