// Function to exponentiate 

// Problem: Given a number and an integer power, return the number raised to its power.

// For example:

//     (power 3 2) = 3^2 = 9

// LISP Code:
    
//     (defun power (x n)
//       "Power raises x to the nth power. N must be an integer >= 0.
// 	   This executes in log n time, because of the check for even n."
// 	   (cond (= n 0) 1)
// 	       ((evenp n) (expt (power x (/ n 2)) 2))
// 		   (t (* x (power x (- n 1)))))

package main

import (
	"fmt"
	"os"
	"strconv"
)

func power(x, n int) int {
	if n == 0 {
		return 1
	}
	if n % 2 == 0 {
		return power(x, n / 2) * power(x, n / 2)
	}

	return x * power(x, n - 1)
}

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])

	fmt.Println(power(x, n))
}