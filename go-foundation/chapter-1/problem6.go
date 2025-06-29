package main

import "fmt"

func main() {
	// format a number with leading zero
	fmt.Printf("Number: %5d\n",78) // 3 blank space + 2 digit = 5
	fmt.Printf("Number: %05d\n",78) // 3 zeros + 2 digit
	fmt.Printf("Number: %o5d\n",78) // octal of 78 + 5d as it is
	fmt.Printf("Number: %05f\n",38.64) //2 digit + 3 ending zero


}