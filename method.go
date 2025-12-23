package main

import "fmt"

type camera struct {
}

func (c *camera) takeAPicture() string {
	return "click"
}

type phone struct {
}

func (p *phone) Ringtone() string {
	return "ring ring"
}

type smartphone struct {
	camera
	phone
}

func mainMethod() {
	output := new(smartphone)
	fmt.Println(output)
	fmt.Println(output.takeAPicture())
	fmt.Println(output.Ringtone())
}