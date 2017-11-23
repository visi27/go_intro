package num

import "math"

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func GetNextPrime(currentNum int) int {
	if IsPrime(currentNum+1) {
		return currentNum +1
	} else {
		return GetNextPrime(currentNum+1)
	}
}