package num

import (
	"math"
)

func CalculatePi(precision int) float64 {
	pi := float64(0)

	for k := float64(0); k < float64(precision); k++ {
		pi += 1 / math.Pow(16, k) * (4/(8*k+1) - 2/(8*k+4) - 1/(8*k+5) - 1/(8*k+6))
	}

	return pi
}
