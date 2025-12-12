package main

import (
	"fmt"
	"reflect"
	"strings"
)

func struct1() {

	type myStruct struct {
		first  int
		second float32
		third  string
	}

assign := new(myStruct)
	assign.first = 10
	assign.second = 12.4
	assign.third = "hello"

	fmt.Println(assign.first)
	fmt.Println(assign.second)
	fmt.Println(assign.third)
	fmt.Println(assign)
}


type person struct{
	firstname  string
	lastName string
}

func fName (p *person){
	p.firstname = strings.ToUpper(p.firstname)
	p.lastName = strings.ToUpper(p.lastName)
}

func fnameMain (){
	var fn person
	fn.firstname = "japheth"
	fn.lastName = "Daniel"
	fName(&fn)
	fmt.Println(fn)

	full := new(person)
	full.firstname = "olawale"
	full.lastName = "glory"
	fName(full)
	fmt.Println(full)

	name3 := &person{"Biba", "wonder"}
	fName(name3)
	fmt.Printf("The name of the person is %s %s\n", name3.firstname, name3.lastName)
}


type reFlect struct{
field1 bool "Yes or No"
field2 string "A great day"
field3 int "Liquidddd"
}

func reflectTag(tt reFlect, ix int){
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)

}

func mainTag(){
	tt:= reFlect{true, "Unknown Anonymous", 1}
	for i :=0; i<3; i++{
		reflectTag(tt, i)
	}

}