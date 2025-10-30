package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	concatResult := concatFunc("Hello", "World")
	fmt.Println(concatResult)

	// call by value demo
	a := 5
	b := 6
	fmt.Println("Call by Value: ")
	fmt.Println("Before swap: ", a, " ", b)
	callByValueSwapExample(a, b)           // arguments a and b. values are copied into func params a and b
	fmt.Println("After swap: ", a, " ", b) // thats why it didnt changed the argument variables values.
	fmt.Println()

	// call by reference demo
	fmt.Println("Call by Reference: ")
	fmt.Println("Before swap: ", a, " ", b)
	callByRefSwapExample(&a, &b)           // arguments a and b. values are copied into func params a and b
	fmt.Println("After swap: ", a, " ", b) // thats why it didnt changed the argument variables values.
	fmt.Println()

	// Variable number of arguments
	variadic := variadicFunc1(1, 2, 3, 4, 5)
	fmt.Println("Variable number of arguments for variadicFunc1(1, 2, 3, 4, 5) -> ", variadic)

	variadicFunc2Sum, variadicFunc2Hello := variadicFunc2("Hello", 1, 2, 3, 4)
	fmt.Println(`variadicFunc2("func variadicFunc2(hello string, nums ...int) (int, string)", 1, 2, 3, 4) -> returns `, variadicFunc2Sum, variadicFunc2Hello)

	// Anonymous functions in golang
	func() {
		fmt.Println("Anonymous function 1")
	}()

	// Assigning to a Variable
	value := func() {
		fmt.Println("Anonymous function 2")
	}
	value()

	// Passing Arguments

	func(print string) {
		fmt.Println(print)
	}("Print")

	// passing anonymous function to another function
	value2 := func(a, b int) int {
		return a + b
	}
	passAnonymousFunctionAsArgument(value2)

	// return an anonymous function from another function.

	fmt.Println("return an anonymous function from another function.")
	valueAnonymous := anonymous()
	fmt.Println(valueAnonymous("Hello", " Anonymous Functions"))
}

func anonymous() func(p, q string) string {
	functionScopeAnonymous := func(p, q string) string {
		return p + q
	}
	return functionScopeAnonymous
}
func passAnonymousFunctionAsArgument(an func(a, b int) int) {
	fmt.Println("Passed anonymous function as argument to parameter ", an(5, 5))
}
func variadicFunc1(nums ...int) int {
	total := 0

	for _, val := range nums {
		total += val
	}
	return total
}
func variadicFunc2(hello string, nums ...int) (int, string) {
	total := 0

	for _, val := range nums {
		total += val
	}
	return total, hello
}
func callByValueSwapExample(a, b int) {
	temp := a
	a = b
	b = temp
	fmt.Println("Call by value -> a: ", a, " b: ", b)
}
func callByRefSwapExample(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
	fmt.Println("Call by Reference -> a: ", *a, " b: ", *b)
}
func concatFunc(str1, str2 string) string {
	return str1 + " " + str2
}
