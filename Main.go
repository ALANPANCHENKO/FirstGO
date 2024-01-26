package main

import (
	"fmt"
	"unsafe"
)

const name = "Alan"

func main() {
	hello := fmt.Sprintf("hello %d", name)
	fmt.Println(name)
	p := 127
	_ = p
	flag := true
	fmt.Printf("%T , %v\n", flag, flag)
	fmt.Println(flag)
	fmt.Printf("%T , %v\n", hello, hello)
	fmt.Println(p)

	var num_1 int8 = 15
	var num_2 int8 = 15
	fmt.Println(num_1 + (num_2))
	fmt.Println(unsafe.Sizeof(num_1))
	fmt.Println(unsafe.Sizeof(num_2))
}
