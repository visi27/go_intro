package num

import (
	"math/big"
)

func CalculateE(precision int) float64 {
	e := float64(0)
	var f = big.Int{}

	for k := float64(0); k < float64(precision); k++ {
		e += (2*k+2)/float64(f.MulRange(1, 2*int64(k)+1).Int64())
	}

	return e


}
