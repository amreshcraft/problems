package main

import "fmt"

// first n natural number
func printFirstNOddNaturalNumber(number int) {
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
if(number % 2 == 0){
	 number = number -1
}
	// recursive call
	printFirstNOddNaturalNumber(number - 2)
	fmt.Print(number, ", ")

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNOddNaturalNumber(number)
}
