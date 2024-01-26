package main

import "fmt"

const name2 = "Alan"

func Print(p int) {
	fmt.Println(p)
}

func Sum(p int, h int) int {
	k := p + h
	return k
}

func multiply(p, h int) (int, int) {
	k := p * h
	l := p + h
	return k, l
}

func division(p int, h int) {
	k := p / h
	fmt.Println(k)
}

func multiply_float(p float64, h float64) {
	k := p * h
	fmt.Println(k)
}

func division_float(p float64, h float64) {
	k := p / h
	fmt.Println(k)
}

func main() {
	_ = name2
	p := 8
	h := 9
	sum := Sum(p, h)
	sum_1, mult := multiply(p, h)
	fmt.Println(sum_1, mult)
	fmt.Println(sum)
	multiply(p, h)
	division(p, h)
	y := 8.9
	u := 9.8
	multiply_float(y, u)
	division_float(y, u)

}
