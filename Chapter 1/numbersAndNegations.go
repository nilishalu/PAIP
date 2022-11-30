//Problem Statement: 
//Given a list of elements, return a list consisting of all the numbers
// in the original list and the negation of those numbers.
// For example, function(1 2 3) return (1 -1 2 -2 3 -3)

// LISP Code:
// (defun numbers-and-negations (input)
//   "Given a list, return only the numbers and their negations."
//   (mappend #'number-and-negation input)
// )

// (defun number-and-negation (x)
//   "If x is a number, return a list of x and -x"
//   (if (numberp x)
//     (list x (-x))
//     nil)
// )

//(numbers-and-negations '(testing 1 2 3 test)) => (1 -1 2 -2 3 -3)

package main

import (
	"fmt"
	"unicode"
	"strconv"
)

func isInteger(num string) bool {
	if num[0] == '+' || num[0] == '-' {
       return isInteger(num[1:])
	}
	
	for _, c := range num {
		if unicode.IsDigit(c) == false {
            return false
		}
	}

	return true
}

func numAndNegation(x string) ([]int) {

	if isInteger(x) {
		num, _ := strconv.Atoi(x)
		elements := []int{num, -num}
		return elements
	}
	return []int{}
}

func numsAndNegations(nums []string) ([]int) {
	//Iterative Approach:
	// var numsAndNegs []int
	
	// for _, num := range nums {
	// 	numsAndNegs = append(numsAndNegs, numAndNegation(num)...)
	// }

	// return numsAndNegs

	//Recursive Approach:
	if len(nums) == 0 {
		return ([]int{})
	}

	return append(numAndNegation(nums[0]), numsAndNegations(nums[1:])...)
}

func main() {
	arr := []string{"1", "+2", "-3", "a34"}
	fmt.Println(numsAndNegations(arr))
}
     