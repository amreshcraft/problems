package main

import "fmt"

// first n natural number
func printFirstNNaturalNumber(number int) {
	// negative number
	if number < 1 {
		return
	}

	// base case
	if number == 1 {
		fmt.Print(1, ", ")
		return
	}
	// recursive call
	printFirstNNaturalNumber(number - 1)
	fmt.Print(number, ", ")

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNNaturalNumber(number)
}
