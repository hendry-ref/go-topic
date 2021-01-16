package helper

import "fmt"

func PrintValueType(data ...interface{}) {
	for _, d := range data {
		fmt.Printf("%v => type : %T\n", d, d)
	}

}
