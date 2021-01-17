package main

import "fmt"

func main() {
	functionBasic()
}

func functionBasic() {
	myName := "hendry"
	tryMe("hendry", "ang", 5, 1, 2, 3)
	tryMe ("leon" , "hart" , 10 , []int{5,6,7}...)
	nameChangerByRef(&myName)
	fmt.Println("Caller function name = " + myName)
}

/* syntax :
func <functionName> ( <var>,<var> <type-1>, <var> <type-2> , <var>...<type-3> )
*/
func tryMe(fName, lName string, age int, nums ...int) {
	fmt.Println("Hello", fName+" "+lName, ", you are", age, "years old")

	fmt.Println("you send ...")
	for _, n := range nums {
		fmt.Println("  >", n)
	}
}

// function accepts the pointer, and will modify the original value
// pass-by pointer is efficient. we don't need to pass in entire value of big data, rather than just pass the pointing address
func nameChangerByRef ( name *string ) *string{
	fmt.Println("Original name = " , *name )
	*name = "Modified-" + *name
	fmt.Println("Inside function name = " + *name)
	return name
}
