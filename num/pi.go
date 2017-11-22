package main

import (
	"fmt"
	"math"
)

func main() {
	pi := float64(0)
	precision := 100

	for k := float64(0); k < float64(precision); k++ {
		pi += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
	}

	fmt.Printf("Calculated PI value for precicion %d is %g\n", precision, pi)
}
