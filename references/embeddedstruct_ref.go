package main

import (
	"github.com/hendry-ref/golang-topic/helper"
	"reflect"
)

func main() {
	embeddedStructBasic()
}

type Animal struct {
	Name   string `json:"name" required max:"100"`
	Origin string `json:"origin"`
}

type Bird struct {
	Animal   // embedding (composition, not inheritance)
	SpeedKPH float32
	CanFly   bool
}

func embeddedStructBasic() {
	// #1 - literal syntax
	myBird := Bird{
		Animal:   Animal{Name: "Eagle", Origin: "south america"},
		SpeedKPH: 124,
		CanFly:   true,
	}

	// #2 - declare and assign. Note that we can have access to Animal's attribute
	myBird2 := Bird{}
	myBird2.Name = "Emu"
	myBird2.Origin = "Australia"
	myBird2.SpeedKPH = 40
	myBird2.CanFly = false

	helper.PrintValueType(
		myBird, // myBird still main.Bird , not inheritance (is-a) of Animal. it's composition (has-a)
		myBird2,
	)

	// ability to retrieve the tag from struct using 'reflect'
	theType := reflect.TypeOf(Animal{})
	field, _ := theType.FieldByName("Name")
	helper.PrintValueType(field.Tag)
}
