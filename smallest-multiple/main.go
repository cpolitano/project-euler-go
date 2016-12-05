// 2520 is the smallest number that can be divided by each of the numbers 
// from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by 
// all of the numbers from 1 to 20?

package main

import "fmt"

func main() {
	
	for i := 1; i < 1000000000; i++ {
		x := 0

		if i % 2 != 0 {
			continue
		}

		for j := 1; j < 21; j++ {
			if i % j == 0 {
				if j == 20 {
					fmt.Println("SMALLEST MULTIPLE", i)
					x++
				}
			} else {
				break
			}
		}

		if x > 0 {
			break
		}
	}

}