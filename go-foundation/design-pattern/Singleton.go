package main

import (
	"fmt"
	"sync"
)

// Singleton struct
type Singleton struct {
	Message string
}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating Singleton Instance...")
		instance = &Singleton{Message: "Hello from Singleton!"}
	})
	return instance
}

func main() {
	// Simulating multiple calls
	s1 := GetInstance()
	s2 := GetInstance()

	// Both should point to the same instance
	fmt.Println("s1 message:", s1.Message)
	fmt.Println("s2 message:", s2.Message)
	fmt.Println("Are s1 and s2 same instance?", s1 == s2)
}
