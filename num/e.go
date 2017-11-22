package main

import (
	"fmt"
	"math/big"
)

func main() {
	e := float64(0)
	precision := 30
	var f = big.Int{}

	for k := float64(0); k < float64(precision); k++ {
		e += (2*k+2)/float64(f.MulRange(1, 2*int64(k)+1).Int64())
	}

	fmt.Printf("Last step is %d/%d\n", (2*precision-1+2), f.MulRange(1, 2*int64(precision-1)+1).Int64())

	fmt.Printf("Calculated e value for precicion %d is %g\n", precision, e)
}
