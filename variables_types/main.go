package main

import "fmt"

// jwt := "sjdhsjhdjshdjsdhshd" // this is not allowed outside method or its scope
var jwt = "sjdhsjhdjshdjsdhshd"         // this is allowed outside method or its scope
var jwt1 string = "sjdhsjhdjshdjsdhshd" // this is allowed outside method or its scope
const tokens string = "sewewewewewewe"  // public variable can be used anywhere in project

const aIOTA = iota

func main() {
	// fmt.Println("Hello, World!")
	// variables:

	// integer constants or literals
	a := 02
	fmt.Println(a)
	b := 1_0_2
	fmt.Println(b)

	// floating constants or literals
	fmt.Println()
	fmt.Println("Floating Point Constants")
	c := 1.0e+2
	fmt.Println(c)
	d := 1.1_234
	fmt.Println(d)
	e := 6.67428e-11
	fmt.Printf("%.20f\n", e)

	// variables with var keyword
	fmt.Println()
	fmt.Println("variables with var keyword")
	var userName string = "Vishnu"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName) // %T	a Go-syntax representation of the type of the value

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn) // %t	a Go-syntax representation of the type of the value

	// int
	var smallValue int = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	// float
	var smallFloatValue float32 = 255.90900888
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue) // float32 precisions around

	var largeFloatValue float32 = 255.90900888
	fmt.Println(largeFloatValue)
	fmt.Printf("Variable is of type: %T \n", largeFloatValue) // float64 handles precisions accurately

	// default values
	var otherVariable int
	fmt.Println(otherVariable)
	fmt.Printf("Variables is of type: %T \n", otherVariable)

	var otherVariable1 string
	fmt.Println(otherVariable1)
	fmt.Printf("Variables is of type: %T \n", otherVariable1)

	fmt.Println(tokens)
	fmt.Printf("Variables is of type: %T \n", tokens)

	// rune literals
	var char1 = 'A'
	fmt.Printf("Character value %c \n", char1) // %c for printing character representation
	fmt.Printf("Character type: %T \n", char1)
	fmt.Printf("Character Unicode: %U \n", char1) // %U for printing Unicode representation

	slc := []rune{'A', 'B', 'C', 'D'}
	for i, value := range slc {
		fmt.Printf("Character Value: %c Character type: %T Character Unicode: %U Position: %d", value, value, value, i)
		fmt.Println()
	}

	// constants
	const isBool = true
	fmt.Println(isBool)

	const complex = 1 + 2i
	fmt.Println(complex)

	// IOTA
	fmt.Println("IOTA")
	fmt.Println(aIOTA)
	fmt.Println(aIOTA)

	var hello uint8 = 'A'
	fmt.Println(hello)
}
