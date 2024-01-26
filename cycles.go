package main

import (
	"fmt"
)

func Example(chislo int) bool {
	return chislo < 20
}

func Age(x int) {
	if x < 18 {
		fmt.Println("Вам нет 18 лет")
	}
	if x <= 18 || x >= 10 {
		fmt.Println("Вы молодец!")
	} else {
		fmt.Println("Вы не молодец!")
	}
}

func Operation(x, y, flag int) int {
	if flag == 1 {
		return x + y
	}
	if flag == 2 {
		return x * y
	}
	if flag == 3 {
		return x / y
	}
	return 0
}

func main() {
	first, second := 100, 10
	var op int
	//fmt.Fscan(os.Stdin, &op)
	op = Operation(first, second, 2)
	fmt.Println(op)

	Age(10)

	Kol := 10
	if zadachi := Example(Kol); zadachi == true {
		fmt.Println("Нужно постараться(")
	} else {
		fmt.Println("Так держать!")
	}
}
