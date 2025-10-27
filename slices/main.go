package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices in Golang")

	var fruits = []string{"Apple", "Banana", "Grapes", "Orange"}
	fmt.Printf("%T\n", fruits)

	fruits = append(fruits, "Mango", "Pineapple")
	fmt.Println(fruits)

	// fruits = append(fruits[1:])
	// fmt.Println(fruits)

	fruits = append(fruits[:3])
	fmt.Println(fruits)

	highScores := make([]int, 3)
	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 678
	// highScores[3] = 890 // this will throw an error as the length of the slice is 3
	// because memory is allocated for 3 elements only
	highScores = append(highScores, 890, 1000, 1200) // append() -> here memory is reallocated again to accommodate new elements
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted High Scores:", highScores)

	// slice declaration
	// only differences is we omit length while declaring a slice
	// slice is just pointer to an array
	var slices []int
	slices = append(slices, 1, 2, 3, 4, 5)

	fmt.Println("Slices before deleting elements:", slices)

	// make() slice creation
	slcMake := make([]int, 5)
	slcMake = append(slcMake, 1, 2, 3, 4, 5, 6)
	fmt.Println("Slices created using make():", slcMake)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	subSlice := slice[1:4]
	fmt.Println(subSlice)
	fmt.Println(len(subSlice), cap(subSlice))

	// subSlice with capacity which is optional parameter
	subSliceWithCap := slice[1:4:5]
	fmt.Println("Sub slice with capacity:", subSliceWithCap)
	fmt.Println(len(subSliceWithCap), cap(subSliceWithCap))

	// String Slices
	s := []string{"g", "o", " ", "i", "s", " ", "s", "w", "e", "e", "t"}
	goSubSlice := s[:2]
	fmt.Println(goSubSlice)

	sweetSubSlice := s[3:]
	fmt.Println(sweetSubSlice)

	// best way of copying all items

	copySlice := s[:]
	fmt.Println(copySlice)

	// number slices
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n1 := n[:6]
	n2 := n[3:8]
	n3 := n[4:10]
	fmt.Println(n1, len(n1), cap(n1))
	fmt.Println(n2, len(n2), cap(n2))
	fmt.Println(n3, len(n3), cap(n3))

	n[4] = 15
	fmt.Println(n, len(n), cap(n))
	fmt.Println(n1, len(n1), cap(n1))
	fmt.Println(n2, len(n2), cap(n2))
	fmt.Println(n3, len(n3), cap(n3))

	// Deep value copying of a slice
	// slice1 := []int{1, 2, 3, 4, 5}
	// slice2 := []int{}
	// slice2 = append(slice2, slice1...)
	// fmt.Println("Slice1:", slice1)
	// fmt.Println("Slice2:", slice2)
	// slice1[0] = 100
	// fmt.Println("After modifying Slice1:")
	// fmt.Println("Slice1:", slice1)
	// fmt.Println("Slice2:", slice2)

	// or
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
	slice1[0] = 100
	fmt.Println("After modifying Slice1:")
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
}
