package main

import (
	"fmt"
	"strings"
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

}
