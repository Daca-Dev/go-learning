package main

import (
	"fmt"
	"strings"
)

// GO DOC URL -> https://pkg.go.dev/?utm_source=godoc

func isPalindrome(text string) bool{
	var textReverse string
	lowerText := strings.ToLower(text)
	for i:= len(text) - 1; i >= 0; i-- {
		textReverse += string(lowerText[i])
	}
	
	if textReverse == lowerText {
		return true
	} else {
		return false
	}
}


func main() {
	slice := []string{"Hello", "world", "!!!"}
	// index and value, you can ignore de index with _
	for i, value := range slice {
		fmt.Println(i, value)
	}

	// only the index
	for i := range slice {
		fmt.Println(i)
	}

	/*
	palindrome exercise
	examples:
	1 ana
	2 amor a roma
	*/
	fmt.Println(isPalindrome("Ana"))
	fmt.Println(isPalindrome("amor a roma"))
	fmt.Println(isPalindrome("casas"))

}
