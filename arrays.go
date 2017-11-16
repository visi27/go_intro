package main

import (
	"fmt"
)

func main() {
	sayans := make(map[string]int)
	sayans["Goku"] = 10000

	power, exists := sayans["vegeta"]
	fmt.Println(power, exists)
	power, exists = sayans["Goku"]
	fmt.Println(power, exists)
}