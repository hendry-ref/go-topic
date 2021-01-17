package main

import (
	"fmt"
)

func main() {
	pointerBasic()
	pointerStruct()
}

func pointerBasic() {

	// & -> get the address
	// *<pointer-variable> -> de-referencing pointer-value of variable
	// *<datatype> -> create variable as a pointer
	//    *int -> a pointer to integer
	oriData := 42
	valData := oriData // set valData as copying value of oriData
	valData = 27
	refData := &oriData // set refData as pointer to oriData
	*refData = 50
	fmt.Println("oriData = ", oriData, ", add = ", &oriData)
	fmt.Println("valData = ", valData, ", add = ", &valData)
	fmt.Println("refData = ", *refData, ", add = ", refData)

	// declaring pointer variable *int
	var ori1 int = 42
	var ref1 *int = &ori1
	fmt.Println(ori1, ref1)
	ori1 = 27
	fmt.Println(ori1, ref1)
}

type myStruct struct {
	foo int
	bar bool
}

func pointerStruct() {
	var valMs myStruct
	valMs = myStruct{10, true}
	fmt.Println(valMs)

	var refMs *myStruct
	fmt.Println(refMs) // pointer has default <nil> value if no value
	refMs = &myStruct{10, true}
	fmt.Println(refMs)

	var refMs1 *myStruct
	refMs1 = new(myStruct) // use new to create instance of it by no value
	(*refMs1).bar = true   // or refMs1.bar as shortcut . (*refMs1) --> dereferencing refMs1 first then assign value true
	(*refMs1).foo = 12
	fmt.Println(refMs1)        // print memory address
	fmt.Println(*refMs1)       // print value
	fmt.Println((*refMs1).foo) // dereference then get value
	fmt.Println(refMs1.foo)    //shorthand for (*refMs1).foo
}
