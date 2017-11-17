package main

import (
	"fmt"
)

type Add func(a int, b int) int

func main() {
	fmt.Println(exec(func(a int, b int) int {
		return a + b
	}))
}
func exec(adder Add) int {
	return adder(1, 2)
}
