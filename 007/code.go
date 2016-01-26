//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

//What is the 10 001st prime number?

package main

import "fmt"

func main() {
	counter := 0
	number := 2
	for {
		if isPrime(number) {
			counter++
		}
		if counter == 10001 {
			fmt.Println(number)
			return
		}
		number++
	}
}

func isPrime(number int) bool {
	counter := 2
	for counter < number {
		if number%counter == 0 {
			return false
		}
		counter++
	}
	return true
}
