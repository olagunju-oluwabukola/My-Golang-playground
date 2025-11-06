package main

import "fmt"

func stringLen() {
	str:= "for loop test"
	fmt.Printf("The string has a length of %d\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("character at position %d is : %c\n", i, str[i])
	}

}

func test2(){
	str2:= "I got the first one"
	fmt.Printf("this string has a length of %d\n", len(str2))
	 for ix:=0; ix < len(str2); ix++{
		fmt.Printf("The character at %d is : %c\n", ix, str2[ix])
	}
}