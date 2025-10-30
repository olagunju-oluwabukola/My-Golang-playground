package main

import (
	"fmt"
	"strings"
)

func userReg() {

	var FName string
	var LName string
	var age int
	users := []string{}
for {
	fmt.Println("Your first name:")
	fmt.Scan(&FName)

	fmt.Println("Your Last Name:")
	fmt.Scan(&LName)

	fmt.Println("How old are you?")
	fmt.Scan(&age)
	users = append(users, FName + " " + LName)


	firstNames := []string{}

	for _, regUser := range users{
		var names = strings.Fields(regUser)
		var userFirstname = names[0]
		firstNames = append(firstNames,userFirstname )
	}
	// fmt.Println("Our registered users are %v\n",firstNames )
	fmt.Printf("The type is %T",firstNames)
	fmt.Printf("The type is %v",len(firstNames))
	}

}