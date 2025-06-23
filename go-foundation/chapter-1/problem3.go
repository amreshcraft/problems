package main

import "fmt"

func main() {
	const PI = 3.141592
	fmt.Printf("PI : %d\n", PI)
	fmt.Printf("PI : %e\n", PI)
	fmt.Printf("PI : %.2f\n", PI)
	// float to string
	str := fmt.Sprintf("%s", PI)
	fmt.Println("\nstr : ", str)
	FloatToString(123.456, 'f', 2)
}

func FloatToString(f float64, spec byte, prec int) string {
	format := fmt.Sprintf("%%.%d%c", prec, spec)
	str := fmt.Sprintf(format, f)
	return str
}
