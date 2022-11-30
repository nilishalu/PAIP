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