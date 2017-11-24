package num

func CalculateFactorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return CalculateFactorial(num-1) * num
	}
}