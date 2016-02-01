//A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

//A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

//As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

//Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

package main

import (
	"fmt"
)

func main() {
	abundants := abundantNumbers() // [1 2 4 7 14]
	sums := allSums(abundants)
	fmt.Println(len(sums))
	fmt.Println(sums)
	var result []int
	for i := 1; i < 28123; i++ {
		if !include(i, sums) {
			result = append(result, i)
		}
	}
	s := 0
	for _, v := range result {
		s += v
	}
	fmt.Println(s)
}

func integers() []int {
	var result []int
	for i := 1; i < 28123; i++ {
		result = append(result, i)
	}
	return result
}

func include(number int, numbers []int) bool {
	for _, v := range numbers {
		if number == v {
			return true
		}
	}
	return false
}

func allSums(numbers []int) []int {
	var result []int
	for i := 1; i < len(numbers); i++ {
		for j := i + 1; j <= len(numbers); j++ {
			//fmt.Printf("i: %d, j: %d\n", i, j)
			n := numbers[i] + numbers[j-1]
			if !include(n, result) {
				result = append(result, n)
			}
		}
	}
	return result
}

func abundantNumbers() []int {
	var result []int
	for i := 12; i <= 28123; i++ {
		//for i := 12; i <= 280; i++ {
		if sumSlice(properDivisors(i)) > i {
			result = append(result, i)
		}
	}
	return result
}

func sumSlice(numbers []int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

func properDivisors(number int) []int {
	var result []int
	for i := 1; i < number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
