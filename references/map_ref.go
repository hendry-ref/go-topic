package main

import (
	"fmt"
	"github.com/hendry-ref/golang-topic/helper"
)

func mapBasic() {
	// map is K:V pair, as in python's dictionary. it's unordered, mutable sequence
	// in map, Key must be comparable data type (eq. string, int, float, bool, struct, array)
	// non-comparable data type ---> slices, maps and other func CANNOT be KEY of the map
	// map is like slice (reference type)
	// #1 - literal approach
	statePop1 := map[string]int{
		"California": 11223344,
		"Texas":      12918182,
		"New York":   99210212,
	}

	// map with value of slice (array)
	sampleMap1 := map[string][]int{
		"hendry": {1, 2, 3, 4},
		"john":   {4, 1, 2},
		"robert": {43, 1, 2},
	}

	// invalidMap := map[[]int]string {}
	// validMap := map[[3]int]string{}

	// #2 - using make(...) and assign
	statePop2 := make(map[string]int)
	statePop2 = map[string]int{
		"California": 11223344,
		"Texas":      12918182,
		"New York":   99210212,
	}

	helper.PrintValueType(
		statePop1,
		sampleMap1,
		statePop2,
		len(statePop2),
	)

}

func mapAccess() {
	fmt.Println("\n=== Map Access ===")
	statePop1 := map[string]int{
		"California": 11223344,
		"Texas":      12918182,
		"New York":   99210212,
	}
	helper.PrintValueType(statePop1["California"])

	// accessing non-exist KEY in map, will return the default type
	// use val,ok := map[k] for checking
	delete(statePop1, "New York") // delete map by Key

	searchKey := "New York" // case sensitive
	if val, ok := statePop1[searchKey]; !ok {
		fmt.Println("Unable to find value for key : ", searchKey)
	} else {
		fmt.Println("Value for key ", searchKey, " = ", val)
	}

}
func main() {
	mapBasic()
	mapAccess()
}
