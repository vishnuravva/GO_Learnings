package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in Golang")

	var list [5]string
	list[0] = "Hello"
	list[1] = "World"
	list[3] = "Arrays"

	fmt.Println(list)
	fmt.Println("Length of array", len(list))

	// ... means “let the compiler figure out the array length automatically”.
	//
	var fruitlist = [...]string{"Apple", "Banana", "Grapes", "Orange", "Test"}
	fmt.Println("Fruit List:", fruitlist)
	fmt.Println("Length of fruit list:", len(fruitlist))

	// If you don’t specify an index, Go automatically assigns the next index.
	// Between index 0 and 3 → nothing was specified.
	// So Go fills those with the zero value of the element type, which for int is 0.
	var numberList = [...]int{1, 3: 3, 4, 5, 6} // {1, 3: 3, 4, 5, 6} → this part is the array composite literal.
	fmt.Println("Number List:", numberList)
	fmt.Println("Length of number list:", len(numberList))
}
