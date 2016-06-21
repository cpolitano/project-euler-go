// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 552 = 3025

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func sumOfSquares(x int) int {
	var sum int
	var i int
	for i <= x {
		sum += i * i
		i++
	}
	return sum
}

func squareOfSums(y int) int {
	var square int
	var i int
	for i <= y {
		square += i
		i++
	}
	return square * square
}

func main() {
	sum := sumOfSquares(100)
	square := squareOfSums(100)
	fmt.Println(square - sum)
}
