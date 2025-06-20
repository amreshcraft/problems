package main

import "fmt"

// first n natural number
func printFirstNEvenNaturalNumberReverse(number int) {
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
	fmt.Print(number, ", ")
	printFirstNEvenNaturalNumberReverse(number - 2)

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNEvenNaturalNumberReverse(number)
}
