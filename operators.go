package main

import (
	"fmt"
	"strconv"
	"strings"
)

func operatorsTest() {
	a := 3
	b := 5

	b = b + a
	fmt.Println( "first instance",b)
	b=5
	b +=a
	fmt.Println( "second instance",b)

	var sentence string = "This is a sentence"

	fmt.Printf(" T/F? Does the sting \"%s\" have prefix  %s?", sentence,"Th")
	fmt.Printf("\n%t\n\n", strings.HasPrefix(sentence, "Th"))

	var userd = "angela is BeaUTiful"
	var lower string
	var upper string

	fmt.Printf("This is in lower case %s\n", lower)
	lower = strings.ToLower(userd)
	fmt.Printf("This is for uppercase : %s\n", upper)
	upper= strings.ToUpper(userd)


	var w int64 = 10
	var r string = "11"
	 xe, _ := strconv.Atoi(r)

	fmt.Println(w+ int64(xe))
}