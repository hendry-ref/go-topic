package main

import (
	"fmt"
	"github.com/hendry-ref/golang-topic/helper"
)

// enumerated constant
const (
	// for quickly assign incremental integer (0...n), we can use iota
	_ = iota // for unassigned (default int=0), we need to assign 'error' to it
	catSpecialist          //iota is scoped to constant block
	dogSpecialist
	snakeSpecialist

	/*  -- OR --
	a = iota
	b = iota
	c = iota
	*/
)

// file size constant
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EZ
	YB
)

// another block of iota constant, will reset its value to 0
const (
	a2 = iota
	b2
	c2
)

const myConst bool = true

func constBasic() {
	const myConst = 10 // inner constant can shadow outer const
	var b float32 = 30.5
	helper.PrintValueType(myConst + b) // myConst can be added with different type.
}

func constEnum() {
	fmt.Println("\n=== Const Enum ===")
	helper.PrintValueType(catSpecialist, dogSpecialist, snakeSpecialist)
	helper.PrintValueType(a2, b2, c2)

	var specialistType int
	helper.PrintValueType(specialistType == catSpecialist)
	specialistType = snakeSpecialist
}

func constFileSize() {
	fileSize := 4000000000.
	helper.PrintValueType(
		KB,
		MB,
		GB,
		TB,
		PB,
		EZ,
		//YB,
		)
	fmt.Printf("%.2fGB\n" , fileSize/GB)
}

func main() {
	constBasic()
	constEnum()
	constFileSize()
}
