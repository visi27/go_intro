package main

import (
	"fmt"
)

var (
	txt2 = "\u6C33\x20brings\x20\x6c\x69\x66\x65."
	txt3 = `\u6C34\x20brings\x20\x6c\x69\x66\x65.`
)


func main() {
	fmt.Println(txt2)
	fmt.Println(txt3)
}