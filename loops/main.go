package main

import "fmt"

func main() {

	fmt.Println("for loop similar to c, java")
	// for loop similar to c, java
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for same as while loop
	fmt.Println("While Loop representation in Golang using for")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++ // if this ignored then it will be infinite loop
	}

	// use of simple range loop -> iterating over strings
	for i, val := range "Hello" {
		fmt.Printf("Index: %v, Character: %c, Unicode: %U ", i, val, val)
		fmt.Println()
	}

	// iterating over a map
	myMap := map[int]string{
		1: "Vishnu",
		2: "Yash",
	}
	for key, val := range myMap {
		fmt.Printf("Key: %v, Value: %v ", key, val)
		fmt.Println()
	}

}
