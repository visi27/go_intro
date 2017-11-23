package num

import (
	"fmt"
	"errors"
)

func PrintPrimalFactors(num int) {
	factors := make([]int, 0)
	fmt.Print(findFactors(num, factors))
}

func findFactors(num int, factors []int) ([]int, error) {
	factor := 2 // First prime factor to test

	if num == 1 {
		return factors, nil
	}

	if (num % factor) == 0 && IsPrime(num/factor) {
		return append(factors, factor, num/factor), nil
	}

	for{
		if factor > 50000 {
			return factors, errors.New("math: can not find primal factor\n")
		}

		if num % factor > 0 || factor > num{
			factor = GetNextPrime(factor)
		} else {
			return findFactors(num/factor, append(factors, factor))
		}
	}
}