package main

import "fmt"

func main2() {
	// number conversion
	num := 20
	bin := fmt.Sprintf("num : %b",num)
	fmt.Println(bin) // 10100
	octal := fmt.Sprintf("num : %O",num) // O - prefix with 0o , o - no prefix
	fmt.Println(octal) // 24
	hex := fmt.Sprintf("num : %x",num)
	fmt.Println(hex) // 14

}