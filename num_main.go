package main

import (
	"hello/num"
	"fmt"
)

func main() {
	n := 10
	num.PrintFibo(n)
	fmt.Printf("\n%d-th Fibo is %d\n", n, num.CalculateNthFibo(10))
}