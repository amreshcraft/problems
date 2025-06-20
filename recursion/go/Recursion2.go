package main

import "fmt"

// first n natural number
func printFirstNNaturalNumberReverse(number int) {
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
	fmt.Print(number, ", ")
	printFirstNNaturalNumberReverse(number - 1)

}
func main() {
	var number int
	fmt.Scan(&number)
	printFirstNNaturalNumberReverse(number)
}
