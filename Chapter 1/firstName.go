// Extracting First Names
//
// Problem: Given the name of a person, return the first name of that person.
//
// For example, if the name is "Captain Jack Sparrow", return "Jack".
// If the name is "Madam General Paula Jones", return "Paula".
// If the name is "Kelly Asbury", return "Kelly".
//
// LISP Code:
//
// 		(defparameter *titles*
// 		  '(Ms Miss Mr Master Dr Mrs Col Colonel Gen General Capt Captain Major Admiral Sir Madam))
//		  "A list of titles that may precede a person's name.")
//
//		(defun first-name (name)
//		  "Select the first name from a name represented as a list"
//		  (if (member (first name) *titles*)
//		      (first-name (rest name))
//		      (first name)))
//


package main

import (
	"fmt"
	"strings"
)

func isATitle(first string, titles []string) bool {
	for _, name := range titles {
		if first == name {
			return true
		}
	}
	return false
}

func firstName(name []string) string {
	titles := []string{"Mr", "Mrs", "Miss", "Ms", "Sir", "Madam", "Dr", "Admiral", "Major", "General"}

	if (isATitle(name[0], titles)) {
		return firstName(name[1:])
	}
	
	return name[0]
}

func main() {
	//fullName := "John Q Public"
	fullName := "Madam Major General Paula Jones"
	name := strings.Split(fullName, " ")
	
    fmt.Println(firstName(name))
}