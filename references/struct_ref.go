package main

import (
	"github.com/hendry-ref/golang-topic/helper"
)

func main() {
	structBasic()
	structAnon()
	structManipulate()
}

// struct is similar to class in Python
// type <Name> struct { <name> <datatype> ... }
type Doctor struct {
	Number     int // Capitalized name to be exported
	actorName  string
	companions []string
}

type DoctorList []struct {
	Number     int
	actorName  string
	companions []string
}

func structBasic() {
	// #1 - recommended way to create instance out of struct
	doctor1 := Doctor{
		Number:    1,
		actorName: "John pertwee",
		companions: []string{
			"one",
			"two",
			"three",
		},
	}

	doctorList := DoctorList{
		{Number: 1, actorName: "A", companions: []string{"1a", "2a", "3a"}},
		{Number: 2, actorName: "B", companions: []string{"1b", "2b", "3b"}},
		{Number: 3, actorName: "C", companions: []string{"1c", "2c", "3c"}},
	}

	// #2 - positional syntax. not recommended (only useful for short-lived/anon struct)
	doctor2 := Doctor{
		2,
		"George Groban",
		[]string{"leo", "nan", "tosh"},
	}
	helper.PrintValueType(doctor1,
		doctor1.actorName,
		doctor1.companions[1],
		doctor2,
		doctorList)
}

func structAnon() {
	anonDoctor1 := struct {
		userid int64
		name   string
	}{1, "hendry"}

	anonList := []struct {
		userid int64
		name   string
	}{
		{1, "hendry"},
		{2, "freecode"},
	}
	helper.PrintValueType(
		anonDoctor1,
		anonList,
		anonList[0],
		anonList[1].name)
}
func structManipulate() {
	// struct is value-type , any copy will be copied it's value rather than reference
	myStruct := struct {
		id   int32
		name string
	}{1, "hendry"}

	myStructVal := myStruct
	myStructVal.name = "lion" //this modification will not modified original value

	myStructRef := &myStruct
	myStructRef.id = 100

	helper.PrintValueType(myStruct, myStructVal, *myStructRef)
}
