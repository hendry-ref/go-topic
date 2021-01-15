package helper

import "fmt"

func PrintValueType(data ...interface{}) {
	for _, d := range data {
		fmt.Printf("%v [%T]\n", d, d)
	}

}
