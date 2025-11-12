package main

import "fmt"

func tCon() {
	var big int64 = 200
	var little int8
	little = int8(big)
	fmt.Println(little)
}