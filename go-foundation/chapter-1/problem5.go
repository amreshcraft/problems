package main

import "fmt"
// right or left alignment
func main() {
	// table - name, age 
	
    fmt.Printf("|%-10s|%-10s|\n", "Name", "Score")
    fmt.Printf("|%-10s|%-10d|\n", "Amresh", 95)
    fmt.Printf("|%-10s|%-10d|\n", "Raj", 88)

	fmt.Println()
	fmt.Printf("|%10s|%10s|\n", "Name", "Age")
    fmt.Printf("|%10s|%10d|\n", "Amresh", 25)
    fmt.Printf("|%10s|%10d|\n", "Raj", 29)



}