package main

import "fmt"

// first n natural number
func printFirstNEvenNaturalNumber(number int) {
	// negative number
	if number <= 1 {
		return
	}

	// base case
	if number == 2 {
		fmt.Print(2, ", ")
		return
	}

	// start with even 
if(number & 1 >= 1){
	 number = number -1
}
	// recursive call
	printFirstNEvenNaturalNumber(number - 2)
	fmt.Print(number, ", ")

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNEvenNaturalNumber(number)
}
