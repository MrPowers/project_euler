//The sum of the squares of the first ten natural numbers is,

//12 + 22 + ... + 102 = 385
//The square of the sum of the first ten natural numbers is,

//(1 + 2 + ... + 10)2 = 552 = 3025
//Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "fmt"

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}

func sumOfSquares(max int) int {
	counter := 1
	result := 0
	for counter <= max {
		result += counter * counter
		counter++
	}
	return result
}

func squareOfSum(max int) int {
	counter := 1
	sum := 0
	for counter <= max {
		sum += counter
		counter++
	}
	return sum * sum
}
