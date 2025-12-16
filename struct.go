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

type T struct {
  a int "This is a tag"
  b int `A raw string tag`
  c int `key1:"value1" key2:"value2"`
}

func mainT(){
	t := T{}
	fmt.Println(reflect.TypeOf(t).Field(0).Tag)
}


type innerS struct{
	in1 int
	in2 int
	string
}

type outerS struct{
	out1 int
	out2 int
	int
	float32
	innerS
}

func outerMain(){
	inner := new(innerS)
	inner.in1 = 10
	inner. in2 = 50
	inner.string = "ciao"
	outer := new(outerS)
	outer.out1 = 10
	outer.out2 = 12
	outer.int = 17
	outer. in1 = 11
	outer.in2  = 18
	outer.string = "hekllo"
	outer.float32 = 23.6

	fmt.Println("outer is :", outer)
	fmt.Println(inner)
}

//anonymous struct
func anonymousStruct (){
	structName := struct{
	firstNmae, lastName string
	} {"barak", "Obama"}
	fmt.Println(structName)

	var structName2 struct{
	firstName, lastName string
	}

	structName2.firstName, structName2.lastName = "Elon", "mussk"
	fmt.Println(structName2)


}


func config (){
	serverConfig := struct{
		HostName string
		Ip string
		Port int
		Environment string
		Credentials struct{
			Username string
			Password string
		}
	}{
		HostName: "localhost",
		Ip: "189.56.00.1",
		Port: 8080,
		Environment: "production",
		Credentials: struct{Username string; Password string}{
			Username: "oluwabukola",
			Password: "1223",
		},
	}
fmt.Println(serverConfig)

	}
type C struct {
x float32
int
string
}
	func anym(){
		c := C {3.24, 30, "Hello"}
		fmt.Println(c.x, c.int, c.string)
		fmt.Println(c)
		b := new(C)
		b.x = 21.2
		b.int = 10
		b.string = "new func"
		fmt.Println(b)
	}