package main

import (
	"fmt"
	"hello/num"
)

func main() {
	n := 10
	num.PrintFibo(n)
	fmt.Printf(" | %d-th Fibonacci number is %d\n", n, num.CalculateNthFibo(10))

	precision := 10
	fmt.Printf("Calculated PI value for precicion %d is %g\n", precision, num.CalculatePi(precision))
	fmt.Printf("Calculated e value for precicion %d is %g\n", precision, num.CalculateE(precision))
}
