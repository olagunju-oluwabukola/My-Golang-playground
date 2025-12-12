package main

import (
	"fmt"
	"regexp"
)

//removes hyphen from the numbers
func RemoveHyphens(text string) string {
	re := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
	updatedText := re.ReplaceAllStringFunc(text, func(match string) string{
		return regexp.MustCompile(`-`).ReplaceAllString(match, "")

	})
	return updatedText
}

func 	RemoveHyphenMain(){
	text:= "John: 123-456-7890 William: 111-222-3333 Steve: 444-555-6666 Thomas: 777-888-9999"
	updatedText := RemoveHyphens(text)
	fmt.Println("Text with hyphens removed from phone numbers:")
	fmt.Println(updatedText)
}