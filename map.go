package main

import (
	"fmt"
	"math/rand"
)

func mapMain() {
	map1 := map[string]int{"one": 1, "two": 2}
	var map3 map[string]int
	map2 := make(map[string]int)
	map3 = map1
	map2["key1"] = 10
	map2["key2"] = 12
	map3["one"] = 12
	map1["three"] = 12

	fmt.Printf("map literal at \"one\" is %d\n :", map1["one"] )
	fmt.Printf("map literal at \"two\" is %d\n :", map1["two"])
	fmt.Printf("map literal at \"three\" is %d\n :", map1["three"])
	fmt.Printf("map literal at \"key 1\" is %d\n :", map2["key2"])
	fmt.Printf("map literal at \"key 2\" is %d\n :", map2["key1"])
	fmt.Printf("map literal at \"key 3\" is %d\n :", map2["key3"])
	fmt.Printf("map literal at \"three\" is %d\n :", map3["three"])


}

func mapCheck(){
	var value string
	var isPresent bool
	var countryCheck = make(map[string]string)
	countryCheck["Nigeria"] = "Abuja"
	countryCheck ["Ghana"] = "Accra"
	countryCheck["Egypt"] = "Cairo"
	value, isPresent = countryCheck["Nigeria"]
	 if isPresent {
		fmt.Println(isPresent, ", the key exist:", value)
	}else{
		fmt.Println(isPresent, "the key does not exist")
	}

	value, isPresent = countryCheck["Europe"]
	if isPresent {
		fmt.Println(isPresent, ", the key exist:", value)
	} else {
		fmt.Println(isPresent, ", the key does not exist", value)
	}

	fmt.Println("before delete:\n", countryCheck)

	delete(countryCheck, "Nigeria")

	fmt.Println("After delete:\n", countryCheck)
}

func randomNumber (){
	for i:=0; i <= 10; i++{
		fmt.Println(i)
	}

	for x := 0; x <=10; x ++{
		y := rand.Intn(20)
		fmt.Println(y)
	}
}

func mapRange (){
	mapFive:= make( map[string]string)

	mapFive ["one"] = "Java"
	mapFive ["two"] = "Kotlin"
	mapFive ["three"] = "django"

	for key, value := range mapFive {
		fmt.Printf( "the key of the map lieral at key one is %s\n and the value is %s\n", key,value)
	}
}


func mapSix (){
	map6 := make(map[string]int)

	map6 ["one"] = 1
	map6 ["two"] = 10
	map6 ["three"] = 4
	map6 ["four"] = 2

	for key, value :=  range map6{
		if value > 4 {
			delete( map6, key)
			fmt.Printf("the key of the deleted map lieral is %v\n and the value is %v\n", key,value)
			fmt.Println(map6)
		} else{
			fmt.Println(map6)
		}
	}
}