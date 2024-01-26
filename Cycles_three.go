package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		if i%4 == 3 {
			if i < 4 {
				continue
			}
			fmt.Println("Число не берется", i)
			continue
		}
		if i%4 == 2 {
			if i < 4 {
				continue
			}
			fmt.Println("Число не берется", i)
			continue
		}
		if i%4 == 1 {
			if i < 4 {
				continue
			}
			fmt.Println("Число не берется", i)
			continue
		}
		fmt.Println(i)
	}
func1:
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 5; j++ {
			if i > 4 {
				break func1
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
}
