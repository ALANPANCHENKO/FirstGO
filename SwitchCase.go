package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Four(x int) int {
	return 4
}

func main() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(4)
	switch value {
	case 3:
		fmt.Println("Value = 3")
	case 1:
		fmt.Println("Value = 1")
	case Four(value):
		fmt.Println("Value = 4")
	default:
		fmt.Println("Не найдено число")

	}
}
