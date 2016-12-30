// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, 
// we can see that the 6th prime is 13.

// What is the 10,001st prime number?

package main

import "fmt"

func main() {

	counter := 0
	number := 2
	primes := make([]int, 0, 8)
	
	for counter < 10001 {
		prime := true

		// check if number is evenly divisible by any previous prime
		for i := 0; i < len(primes); i++ {
			remainder := number % primes[i]
			if remainder == 0 {
				// number is not prime
				number++
				prime = false
				break
			}
		}

		// number is prime
		if counter == 10000 && prime {
			fmt.Println("10001st Prime -", number)
			break
		} else if prime {
			// add prime to slice of primes
			primes = append(primes, number)
			counter++
			number++
		}
	}

}

