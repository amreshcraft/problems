package go
package main

import "fmt"

// first n natural number
func printFirstNOddNaturalNumberReverse(number int) {
	// negative number
	if number < 1 {
		return
	}

	// base case
	if number == 1 {
		fmt.Print(number, ", ")
		return
	}

	// start with Odd 
if(number % 2 == 0 ){
	 number = number - 1
}
	// recursive call
	fmt.Print(number, ", ")
	printFirstNOddNaturalNumberReverse(number - 2)

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNOddNaturalNumberReverse(number)
}
