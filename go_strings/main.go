package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("Welcome to Strings in Golang")
	stringName := "Hello World \"Golang\""
	fmt.Println(stringName)
	stringName1 := `Hello World "Golang"`
	fmt.Println(stringName1)

	str := "Vishnu"
	strFind := "V"
	fmt.Println("Index of V:", strings.Index(str, strFind))

	interpretedStr := "Here \n escape sequences are allowed"
	fmt.Println(interpretedStr)
	rawStr := `It treats everything \n as raw \n also escape sequences`
	fmt.Println(rawStr)

	interpretedStr1 := "\n"
	rawStr1 := `\n`
	fmt.Println([]byte(interpretedStr1)) // [10] \n is treated as single character new line here
	fmt.Println([]byte(rawStr1))         // [92 110] treat each character as byte like slash and n as different

	// str with emoji
	strEmoji := "Hello 🌍"
	fmt.Println("Length of strEmoji: ", len(strEmoji))
	fmt.Println("strEMoji internal bytes representation ", []byte(strEmoji))

	var data rune = '👍'
	fmt.Println("Data: ", data)                         // prints 128077 which is unicode code point
	fmt.Println("Data as byte: ", []byte(string(data))) // prints [240 159 152 175] which is the UTF-8 encoding
	fmt.Printf("Data of rune %c", data)

	// Creating and initializing a string
	// using shorthand declaration
	mystr := "Welcome to Learning Golang"

	// looping over strings
	for _, val := range mystr {
		fmt.Printf("Value of character: %v, Character: %c", val, val)
		fmt.Println()
	}

	//  accessing the individual byte of the string
	myStr1 := "Welcome to GeeksforGeeks"
	myStrBytes := []byte(myStr1)
	fmt.Println(myStrBytes) /// prints each byte of the string in decimal format

	// getting bytes in hexadecimal format
	for _, val := range myStr1 {
		fmt.Printf("Character: %c, Value: %x", val, val)
		fmt.Println()
	}

	// create a string form the slice of bytes and runes
	myslice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	strFromSlice := string(myslice1)
	fmt.Println("String from slice: ", strFromSlice)

	myslice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	strFromSlice2 := []rune(myslice2)
	fmt.Println("String from slice2: ", string(strFromSlice2))

	// Finding len of string
	lenStr := "Hello👋"
	fmt.Println("Length of myStr using len():", len(lenStr))                       // return number of bytes -> 9
	fmt.Println("Length of myStr using each rune", utf8.RuneCountInString(lenStr)) // return number of runes (each character) -> 6

	// using string functions from strings package
	strFuncTest := "@@Hello World!!"
	fmt.Println(strings.Trim(strFuncTest, "@!"))

	// splitting a string
	splitStr := "Hello World from Golang"
	fmt.Println(strings.Split(splitStr, " ")) // returns a slice of strings

	// s := "Welcome,to,GeeksforGeeks"
	// fmt.Println("", s)

	// result := strings.SplitN(s, ",", 2)
	// fmt.Println("Result:", result)
	s := "Welcome,to,GeeksforGeeks"
	fmt.Println("", s)

	result := strings.SplitAfterN(s, ",", 2)
	fmt.Println("Result using SplitN:", result)

	// comparing strings
	strComp1 := "Golang"
	strComp2 := "Golang1"
	// strComp3 := "GoLang"
	fmt.Println("strComp1 == strComp2 ", strComp1 == strComp2)
	fmt.Println("strComp1 < strComp2 ", strComp1 < strComp2)

}
