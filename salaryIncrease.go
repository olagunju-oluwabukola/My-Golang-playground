package main

import "fmt"

type employee struct {
	salary float32
}

func (e *employee) giveRaise(pct float32) {
	e.salary += e.salary * pct
}

func salaryMain() {
	var s = new(employee)
	s.salary = 100000
	s.giveRaise(0.04)
	fmt.Printf("employee noe makes %f :", s.salary)
}