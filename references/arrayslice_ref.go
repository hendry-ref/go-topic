package main

import (
	"fmt"
	"github.com/hendry-ref/golang-topic/helper"
)

func main() {
	arrayBasic()
	array2D()
	arrayCopy()
	sliceCopy()
	sliceSlicing()
	sliceModification()
}

func arrayBasic() {
	// array is fixed size, mutable and only can store 1 data type
	// <name> := [<size>]<type>{<data>}
	theArrFixed := [5]int{10, 20, 30, 40}
	theArrDynamic := [...]string{"hendry", "robert", "john"}

	// slice is dynamically generated size
	// slice has both length (actual current data) and capacity (the total capacity)
	// <name> := []<type>{<data>}
	theSlice := []float64{3.14, 1.23, 4.67}
	helper.PrintValueType(
		theArrFixed, len(theArrFixed),
		theArrDynamic, theArrDynamic[0], len(theArrDynamic),
		theSlice, len(theSlice), cap(theSlice),
	)
}

func sliceModification() {
	// all modification done to slice, will be reflected in both modified/original slices
	// no other way to solve this, other than using loop to manually "copy" original slice
	fmt.Println("\n=== Slice Modification ===")

	// appending the slice
	theSlice := []float64{3.14, 1.23, 4.67}
	theSlice = append(theSlice, 5.55, 6.66, 7.77, 8.88)
	theSlice = append(theSlice, []float64{9.99, 10.10, 11.11}...) // spreading

	helper.PrintValueType(
		theSlice)

	theS := []int{0, 1, 2, 3, 4, 5, 6}
	thePop := theS[:len(theS)-1]            // pop (remove last)
	theShift := theS[1:]                    // shift (remove first)
	theMid := append(theS[:3], theS[5:]...) // to remove middle (3 and 4)

	helper.PrintValueType(theS,
		thePop,
		theShift,
		theMid,
	)

}
func sliceSlicing() {

	fmt.Println("\n=== Slice Slicing ===")
	theSlice := []float64{1.11, 2.22, 3.33, 4.44, 5.55, 6.66}
	helper.PrintValueType(
		theSlice[:],   // slice of all
		theSlice[1:4], // from index=1 to index=3 (4-1)
		theSlice[:3],  // from index=0 to index=2 (3-1)
		theSlice[2:],  // from index=2 to end
	)
}
func array2D() {
	fmt.Println("\n=== Array 2D ===")
	// #1 -> create and assign
	var theMatrix1 [3][3]int
	theMatrix1[0] = [3]int{1, 2, 3}
	theMatrix1[1] = [3]int{4, 5, 6}
	theMatrix1[2] = [3]int{7, 8, 9}

	// #2 -> direct assign
	theMatrix2 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// #3 -> using var
	var theMatrix3 = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	helper.PrintValueType(theMatrix1, theMatrix2, theMatrix3)
}

func arrayCopy() {
	fmt.Println("\n=== Array Copy ===")

	// Array is pointing to the values , not underlying address
	// Copying array will copy its value not the actual address
	originalArr := [...]int{1, 2, 3}
	copyValueArr := originalArr
	copyValueArr[1] = 100 // modifying the copy will not modify original array
	helper.PrintValueType(
		originalArr,
		copyValueArr,
	)

	// Copying array by reference ( pointing to the address with & )
	copyRefArr := &originalArr
	copyRefArr[1] = 500
	helper.PrintValueType(
		originalArr,
		copyRefArr,  // this is pointing to address
		*copyRefArr, // use * to de-reference it
	)
}

func sliceCopy() {
	fmt.Println("\n=== Slice Copy ===")

	// Slice by default is reference type (already pointing to the address, not value)
	// Copying slice will copy by reference, changes to copied slice will also change the original slice
	originalSlice := []int{1, 2, 3}
	copySlice := originalSlice
	copySlice[1] = 100
	helper.PrintValueType(
		originalSlice,
		copySlice)
}
