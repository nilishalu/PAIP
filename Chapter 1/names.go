package main

import (
	"fmt"
	"strings"
)

func lastName(name string) string {
	last_name := strings.Split(name, " ")
	return last_name[len(last_name) - 1]
}

func firstName(name string) string {
	first_name := strings.Split(name, " ")
	return first_name[0]
}

func main() {
	fmt.Println(firstName("Shiva Dasari"))
	fmt.Println(firstName("John Q Public"))
	fmt.Println(firstName("Real Admiral Grace Murray Hopper"))
	fmt.Println(firstName("Spot"))
	fmt.Println(firstName("Aristotle"))

	
	fmt.Println(lastName("Shiva Dasari"))
	fmt.Println(lastName("John Q Public"))
	fmt.Println(lastName("Real Admiral Grace Murray Hopper"))
	fmt.Println(lastName("Spot"))
	fmt.Println(lastName("Aristotle"))
}