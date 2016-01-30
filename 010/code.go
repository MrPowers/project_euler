//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

//Find the sum of all the primes below two million.

package main

import "fmt"
import "math"

func main() {
	primes := primeGenerator(2000000)
	sum := 0
	for _, num := range primes {
		sum += num
	}
	fmt.Println(sum)
}

func primeGenerator(max int) []int {
	var result []int
	for i := 2; i < max; i++ {
		if isPrime(float64(i)) {
			result = append(result, i)
		}
	}
	return result
}

func isPrime(number float64) bool {
	upper := math.Sqrt(number)
	for i := float64(2); i <= upper; i++ {
		if int(number)%int(i) == 0 {
			return false
		}
	}
	return true
}
