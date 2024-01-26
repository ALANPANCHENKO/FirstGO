package main

import "fmt"

var multiplier func(x, y int) int

func sum(x, y int, action func(x, y int) int) int {
	fmt.Println(x, y, action(x, y))

	return action(x, y) + x + y
}

func someFunc(vari int) func(x int) int {
	first_func := func(x int) int { return x / vari }
	return first_func
}

func main() {
	first, second := 1, 2

	//multiplier = func(x, y int) int { return x * y }
	//fmt.Println(multiplier(first, second))
	//
	//multip := func(x, y int) int { return y / x }
	//fmt.Println(multip(first, second))

	sumFunc := func(x, y int) int { return x + y }
	fmt.Println(sumFunc(first, second))
	del := func(x, y int) int { return y / x }
	_ = del
	//fmt.Println(sum(first, second, sumFunc))
	//fmt.Println(sum(first, second, del))

	perem := someFunc(10)
	_ = perem
	//fmt.Println(perem(10))
}
