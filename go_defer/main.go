package main

import "fmt"

func main() {
	// fmt.Println("Start")
	// defer fmt.Println("End")
	// fmt.Println("Middle")

	var x string = "sdsd"
	defer fmt.Println("Deferred x value: ", x)
	x = "10"
	fmt.Println("x value changed to ", x)
}
