// Extracting First Names
//
// Problem: Given the name of a person, return the last name of that person.
//
// For example, if the name is "Captain Jack Sparrow", return "Sparrow".
// If the name is "Madam General Paula Jones", return "Jones".
// If the name is "Kelly Asbury, Jr", return "Asburry".
//
// LISP Code:
//
// 		(defparameter *titles*
// 		  '(MD MD, Jr Jr,))
//		  "A list of titles that may precede a person's name.")
//
//		(defun last-name (name)
//		  "Select the last name from a name represented as a list"
//		  (if (member (last name) *titles*)
//		      (last-name (head name))
//		      (last name)))
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

func lastName(name []string) string {
	titles := []string{"MD,", "Jr", "MD", "Jr,"}
    
	last := len(name) - 1
	if (isATitle(name[last], titles)) {
		return lastName(name[:last])
	}

	last_name := name[len(name) - 1]
    length := len(last_name)

	if (last_name[length - 1] == ',') {
		return last_name[:length - 1]
	}
	
	return name[len(name) - 1]
}

func main() {
	//fullName := "John Q Public"
	// fullName := "Madam Major General Paula Jones Jr"
	//fullName := "Morton Downey, Jr"
	name := strings.Split("Rex Morgan MD,", " ")
	
    fmt.Println(lastName(name))
}