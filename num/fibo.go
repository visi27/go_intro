package num

import "fmt"

func PrintFibo (n int) {
	f1, f2 := 1, 1
	var tmp int;
	fmt.Printf("%d %d ", f1, f2)

	for i :=3; i<=n; i++ {
		tmp = f1 + f2
		f1 = f2
		f2 = tmp
		fmt.Printf("%d ", f2)
	}
}

func CalculateNthFibo(n int) int {
	if n==0 || n==1 {
		return 1
	}

	return CalculateNthFibo(n - 1) + n
}