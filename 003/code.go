//The prime factors of 13195 are 5, 7, 13 and 29.

//What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func main() {
	factors := primeFactors(600851475143)
	fmt.Println(factors)
	fmt.Println(max(factors))
}

func firstPrimeFactor(number int) int {
	n := 2
	for n < number {
		if number%n == 0 {
			return n
		}
		n++
	}
	return number
}

func primeFactors(number int) []int {
	var result []int
	for {
		factor := firstPrimeFactor(number)
		result = append(result, factor)
		if number == factor {
			return result
		}
		number = number / factor
	}
}

func max(numbers []int) int {
	result := numbers[0]
	for _, value := range numbers {
		if value > result {
			result = value
		}
	}
	return result
}
