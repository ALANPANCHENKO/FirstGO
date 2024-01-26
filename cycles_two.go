package main

import "fmt"

func Computer(x int) int {
	fmt.Println("Проверка дисков")
	return x
}

func main() {
	i := 0
	i++
	fmt.Println(i)
	j := 10
	j--
	fmt.Println(j)
	x := 500
	if memory := Computer(x); memory > 256 {
		fmt.Println("Память переполнена.Очистка памяти...")
		for i = memory; i >= 256; i-- {
			memory = i
		}
		fmt.Println(memory)
		fmt.Println("Очитска выполнена")
	}
	for i = 1; i < 10; i++ {
		for j = 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
	var flag = true
	for flag == true {
		meaning := 0
		for i = meaning; i <= 10; i++ {
			meaning = i
			if meaning == 10 {
				flag = false
				fmt.Println("Конец программы")
			}
		}
	}
}
