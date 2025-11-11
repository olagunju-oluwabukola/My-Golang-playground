package main

import "fmt"

func vs(s string) int {
	count := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}
return count
}

func vowelCount() {
	input := "Hello world"
	fmt.Println(vs(input))
}
