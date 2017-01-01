// The Fibonacci sequence is defined by the recurrence relation:

// Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
// Hence the first 12 terms will be:

// F1 = 1
// F2 = 1
// F3 = 2
// F4 = 3
// F5 = 5
// F6 = 8
// F7 = 13
// F8 = 21
// F9 = 34
// F10 = 55
// F11 = 89
// F12 = 144
// The 12th term, F12, is the first term to contain three digits.

// What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

package main

import (
	"fmt";
	// "strconv"; // does not work for numbers > 20 digits
	"math/big"
)

func main() {

	index := 0
	previous_value := big.NewInt(0)
	current_value := big.NewInt(1)

	// creates the smallest number with 1000 digits
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	// loop while current value's digits are less than 1000
	for previous_value.Cmp(&limit) < 0 {
		// next fibonacci number
		previous_value.Add(previous_value, current_value)
		// swap
		previous_value, current_value = current_value, previous_value
		// increment index
		index++
	}
	fmt.Println(index, previous_value) // 1000-digit Fibonacci number
}

