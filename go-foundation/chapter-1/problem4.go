package main

import "fmt"

type Stringer interface {
	String() string
}

type Person struct{
	Name string;
	Age int16;
}

 func (p Person) String() string{
	return fmt.Sprintf("Name: %s, Age: %d",p.Name,p.Age)
}

func main() {
	p:=Person{
		Name:"Amresh",
		Age:25,
	}
	fmt.Println(p)

}