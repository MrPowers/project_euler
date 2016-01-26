// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"
import "strconv"

func main() {
	numbers := threeDigitProducts()
	fmt.Println(max(palindromes(numbers)))
}

func max(numbers []int) (result int) {
	result = numbers[0]
	for _, number := range numbers {
		if number > result {
			result = number
		}
	}
	return
}

func palindromes(numbers []int) []int {
	var result []int
	for _, number := range numbers {
		if isPalindrome(number) == true {
			result = append(result, number)
		}
	}
	return result
}

func isPalindrome(number int) bool {
	s := strconv.Itoa(number)
	return s == Reverse(s)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func threeDigitProducts() []int {
	x := 100
	y := 100
	var result []int
	for x < 1000 {
		y = 100
		for y < 1000 {
			result = append(result, x*y)
			y++
		}
		x++
	}
	return result
}
