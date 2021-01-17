package main

import (
	"fmt"
	"github.com/hendry-ref/golang-topic/helper"
)

func main() {
	datatypeBasic()
	bitWiseOperator()
	stringDatatype()
}

func datatypeBasic() {
	fmt.Println("\n=== datatype basic ===")
	var theBool = true      // default: false
	var theString = "Hello" // default : ""

	// int and int8 cannot be combined in operation (+,-,%,*,/)
	var theInt = 1          // default : 0 , 32 for 32-bit system, 64 for 64-bit system
	var theInt8 int8 = 8    // -128 to 127
	var theInt16 int16 = 16 // -32,768 to 32,767
	var theInt32 int32 = 32 // -2,147,483,648 to 2,147,483,647
	var theInt64 int64 = 64 // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	var theUint uint = 1
	var theUint8 uint8 = 8    // 0 to 255
	var theUint16 uint16 = 16 // 0 to 65,535
	var theUint32 uint32 = 32 // 0 to 4,294,967,295

	// floating number
	var theFloat32 float32 = 32.32 // default : 0
	var theFloat64 = 64.64

	// complex number
	var theComplex64 complex64 = 64 + 64.5i // default : 0+0i
	var theComplex128 = 128 + 128.5i
	var myComplex complex64 = complex(10, 10.5) // real and imag number is provided

	// console printing
	helper.PrintValueType(theBool,
		theString,
		theInt, theInt8, theInt16, theInt32, theInt64,
		theUint, theUint8, theUint16, theUint32,
		theFloat32, theFloat64,
		theComplex64, imag(theComplex64), real(theComplex64),
		theComplex128, imag(theComplex128), real(theComplex128),
		myComplex)
}
func bitWiseOperator() {
	fmt.Println("\n=== bitOperator ===")
	a := 10 // 1010000
	b := 3  // 0011
	helper.PrintValueType(
		a&b,  // 0010 => 2 , & (bitwise AND) : true when both side has bit set
		a|b,  // 1011 => 11 , | (bitwise OR) : true when either side has bit set
		a^b,  // 1001 => 9 , ^ (bitwise XOR) : true when either side has bit set, but not both.
		a&^b, // 1000 => 4 , &^ (bit clear AND NOT) : true when none side has bit set.
		a<<4, // 1010 << 4 becomes 10100000 (notice it's shifted 4 with padded 0)
		a>>1, // 1010 >> 1 becomes 101
	)
}

func stringDatatype() {
	fmt.Println("\n=== string datatype ===")

	// String represents any UTF-8 character
	myStr := "this is a string"
	myStr2 := ", great!"
	helper.PrintValueType(myStr,
		myStr[0],         // will get the byte data. have to be converted back to string
		string(myStr[5]), // convert byte back to string
		myStr+myStr2,     // string concatenate
		[]byte(myStr),    // convert string to slices of byte (byte alias of uint8)
	)

	// Rune represents any UTF-32 character
	var myRune rune = 64 // default : 0
	helper.PrintValueType(
		myRune, // rune is an alias of int32
	)

}
