package main

import (
	"fmt"
	"math"
)

func main() {
	ifElseBasic()
	switchBasic()
	switchCheckType()
	forLoopBasic()
	forLoopBreak()
	forLoopCollections()
}

func ifElseBasic() {
	/* syntax :
	conditional : && || > >= < <= != ! ==
	if <condition> {
		...
	} else if {
		...
	} else {
		...
	}

	if <statement>; <condition> {
		...
	}
	*/

	myNum := 0.123

	// if math.Abs (myNum / math.Pow(math.Sqrt(myNum), 2), 1) < 0.001
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("these are different")
		fmt.Println(myNum, " vs ", math.Pow(math.Sqrt(myNum), 2))
	}
}

func switchBasic() {
	/* syntax:
	- switch does not require "break" statement
	- <item> must be unique across all case
	- we can use fallthrough to continue to next statement. it's logicless, always fallthrough

	switch (<statement>;) <condition> {
	case <item1>:
		...
		(fallthrough)
		(break) -> immediately break out from the switch-case
	case <item2>, <item3>, <item4> :
		...
	default:
		...
	*/
	switch 6 { // or switch i := 2+3; i {
	case 1, 3, 5:
		fmt.Println("in one, three, five")
	case 2, 4, 6:
		fmt.Println("in two, four, six")
	default:
		fmt.Println("Not in the range")
	}

	i := 5 % 3
	switch {
	case i <= 10:
		fmt.Println("result less than 10")
		fmt.Println("  > first case")
		fallthrough
	case i <= 20:
		fmt.Println("result less than 10")
		fmt.Println("  > second case")
	}
}

func switchCheckType() {
	var i interface{} = "hello world"
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case string:
		fmt.Println("i is a string")
	case bool:
		fmt.Println("i is a boolean")
	default:
		fmt.Println("unknown type")
	}
}

func forLoopBasic() {
	// for loop support 'continue' and 'break' keyword as well

	// #1 - standard for <init>; <condition>; <increment> {...}
	for i := 0; i < 5; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	// #2 - assign i value, and we can leave the <init> empty
	i := 0
	for ; i < 5; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println()

	// #3 - have incremental inside, and we can leave the <init> and <increment> empty
	// this is aka While...loop
	x := 0
	for x < 5 { // or for ; x< 5 ; {...}
		fmt.Print(x, ", ")
		x++
	}
	fmt.Println()

	// #4 - have the logical test inside the loop and use "break" to terminate, we can leave everything empty
	y := 0
	for {
		fmt.Print(y, ", ")
		y++
		if y >= 5 {
			break
		}
	}
	fmt.Println()

	// from i = 0 and j = 5 , to i < 5 , increment i by 1 and j by 2
	for i, j := 0, 5; i <= 5; i, j = i+1, j+2 {
		fmt.Print(i, j, " - ")
	}
	fmt.Println()

}

func forLoopBreak() {
Outer: // label for the "break" keyword inside nested loop
	for i := 1; i <= 3; i++ {
		// Inner:  // inner label
		for j := 1; j <= 3; j++ {
			fmt.Printf("%v * %v = %v\n", i, j, i*j)
			if i*j >= 3 {
				// break  // this will only break the "j-loop" (similar with break Inner)
				// break Inner  // this will break out to the Inner: label
				break Outer
			}
		}
	}

}

func forLoopCollections() {
	// looping array
	myArr := [...]int{10, 20, 30}
	fmt.Println("Array : ")
	for k, v := range myArr {
		fmt.Printf(" > k : %v -> v : %v\n", k, v)
	}

	// looping slice
	mySli := []int{10, 20, 30}
	fmt.Println("Slice : ")
	for k, v := range mySli {
		fmt.Printf(" > k : %v -> v : %v\n", k, v)
	}

	// looping map
	myMap := map[int]string{
		10: "ten",
		20: "twenty",
		30: "thirty",
	}
	fmt.Println(" Map : ")
	for k, v := range myMap {
		fmt.Printf(" > k : %v -> v : %v\n", k, v)
	}

	// looping string
	myStr := "Hello world!"
	fmt.Println(" String : ")
	for k, v := range myStr {
		fmt.Printf(" > k : %v -> v : %v\t data : %v\n", k, v, string(v))
	}
}
