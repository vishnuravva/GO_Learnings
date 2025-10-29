package main

import "fmt"

func evenOdd(n int) bool {
	return n%2 == 0
}

func main() {
	var evenOddNo int
	fmt.Println("Enter a number: ")
	// fmt.Scanf("%v", &n)
	fmt.Scanln(&evenOddNo)
	isEven := evenOdd(evenOddNo)

	if isEven {
		fmt.Println("Even Number ", evenOddNo)
	} else {
		fmt.Println("Odd Number ", evenOddNo)
	}
}
