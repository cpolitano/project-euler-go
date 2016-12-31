// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

// Find the sum of all the primes below two million.

package main

import "fmt"

func main() {

	number := 2
	primes := make([]int, 0, 50)
	sum := 0
	
	for number < 2000000 {
		prime := true

		// check if number is evenly divisible by any previous prime
		for i := 0; i < len(primes); i++ {
			remainder := number % primes[i]
			if remainder == 0 {
				// number is not prime
				prime = false
				break
			}
		}

		// number is prime
		if prime {
			// add prime to slice of primes
			primes = append(primes, number)
			sum += number
		}

		number++
	}

	// find sum of all primes below 2 million
	fmt.Println(sum)

}

