package num

import (
	"fmt"
	"unicode/utf8"
)

func BinaryToDecimal(binNum string) int {
	decNum := 0
	binLength := utf8.RuneCountInString(binNum)
	for pos, char := range binNum {
		power := binLength - pos - 1
		decNum += int(char - '0') * (PoW(2, power))
		fmt.Println(int(char - '0'), power)
	}

	return decNum
}

func DecimalToBinary(num int) {
	binaryNum := make([]int, 0)

	reminder := num % 2
	num = num / 2
	for num >= 0 {
		binaryNum = append(binaryNum, reminder)
		if num == 0 {
			break
		}
		reminder = num % 2
		num = num / 2
	}

	for i:=0; i<len(binaryNum); i++ {
		fmt.Print(binaryNum[len(binaryNum)- i - 1])
	}
}

func PoW(a, n int) int {
	var i, result int
	result = 1
	for i = 0; i < n; i++ {
		result *= a
	}
	return result
}