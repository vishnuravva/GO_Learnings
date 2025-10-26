package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Basuc user input handling
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Enter some text")
	input, _ := reader.ReadString('\n') // comma ok syntax
	fmt.Println(input)
	fmt.Printf("Type of input is %T \n", input)

	// More on Comma ok syntax
	// example 1: finding key in map
	myMap := map[string]int{"apple": 5, "banana": 10}
	value, ok := myMap["apple"]
	if ok {
		fmt.Println("Key found")
		fmt.Println("Value is:", value)
	} else {
		fmt.Println("Key not found")
	}

	value1, ok := myMap["cherry"]
	if ok {
		fmt.Println("Key found")
		fmt.Println("Value is:", value1)
	} else {
		fmt.Println("Key not found")
	}

	// Example 2: strconv
	conv, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("Converted value is:", conv)
	} else {
		fmt.Println("Error converting value:", err)
	}

	conv1, err := strconv.Atoi("abc")
	if err == nil {
		fmt.Println("Converted value is:", conv1)
	} else {
		fmt.Println("Error converting value:", err)
	}
}
