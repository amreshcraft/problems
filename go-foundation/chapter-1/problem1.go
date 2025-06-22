package main

import "fmt"

type Person struct {
	Name string
	Age  int16
}

func main1() {
	num := -4.56
	fmt.Printf("Value of num(v) : %v \n", num)
	fmt.Printf("Value of num(+v) : %+v \n", num)
	fmt.Printf("Value of num(#v) : %#v \n", num)
	// ye number m same hi behave krta h - difference to array slice struct string m h

	person := Person{
		Name: "Amresh",
		Age:  25,
	}
	fmt.Printf("Value of Person(v) : %v \n", person)   // only person value
	fmt.Printf("Value of Person(+v) : %+v \n", person) // person key + value
	fmt.Printf("Value of Person(#v) : %#v \n", person) // main.Person{Name:"Amresh", Age:25}

}
