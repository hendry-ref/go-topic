package main

import (
	"fmt"
	"github.com/hendry-ref/golang-topic/helper"
	"strconv"
)

func main() {
	variableDeclaration()
	globalVariable()
	variableConversion()
}

// package-level variable must use "var <name> <type>" format. we can also group them in var(...)
var (
	pkgVar1 string  = "hello"
	pkgVar2 float32 = 3.14
)

func variableDeclaration() {

	// #1 - declare and assign - useful when we need to declare variable that to be assigned later
	var myInt1 int
	myInt1 = 42  // we can only assign (=) but not re-declare (:=)
	helper.PrintValueType(myInt1)

	// #2 - declare and assign inline - useful when we want to explicitly declare the datatype
	var myInt2 int64 = 32
	helper.PrintValueType(myInt2)

	// #3 - implicit declare and assign - quick but unable to explicitly assign the datatype
	myInt3 := 70.5
	helper.PrintValueType(myInt3)

	// #4 - retrieve from package level variable
	helper.PrintValueType(pkgVar1)
	helper.PrintValueType(pkgVar2)

	// #5 - overwrite package-level variable (shadowing)
	pkgVar1 := "super"  // we can both assign (=) and re-declare (:=) for pkg-level var
	helper.PrintValueType(pkgVar1)


}

func globalVariable () {
	// lower-case variable will be scoped in the PACKAGE. all files within the package can access it
	// upper-case variable will be scoped in the GLOBAL (exported)
	// variable in {...} will be scoped in the BLOCK
	var MyGlobalVariable string = "this is globally available"
	fmt.Println(MyGlobalVariable)
}

func variableConversion() {
	var myInt int = 75
	var myFloat float32 = 42.5
	var myStr string = "12345"

	helper.PrintValueType(myInt)

	// string to int
	helper.PrintValueType(string(myInt))  // it'll convert the unicode. use strconv for proper conversion
	helper.PrintValueType(strconv.Itoa(myInt))

	// int to string
	if intConverted, err := strconv.Atoi(myStr); err != nil {
		helper.PrintValueType(intConverted)
	}

	// other conversion
	helper.PrintValueType(rune(myInt))
	helper.PrintValueType(float32(myInt))
	helper.PrintValueType(int32(myFloat))  // losing decimal data

}


