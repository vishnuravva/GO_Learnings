package main

import "fmt"

//	type Address struct {
//		street string
//		city string
//		state string
//		country string
//		pincode int
//	}

// or better way below 👇🏻
type Address struct {
	street, city, state, country string
	pincode                      int
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to Structs in Golang")
	var a Address = Address{"Tagore Nagar", "Mumbai", "Maharashtra", "India", 400017}
	fmt.Println("Address struct: ", a)
	a.street = "Hello"
	fmt.Println("Address struct after changing street: ", a)

	user := User{"Vishnu", "vishnu@gmail.com", true, 22}
	fmt.Println("User struct ", user)
	fmt.Printf("User struct raw %+v ", user)

}
