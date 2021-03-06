//Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

//1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

//By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import "fmt"

func main() {
	numbers := []int{1, 2}
	for numbers[len(numbers)-1] < 4000000 {
		nextNumber := numbers[len(numbers)-1] + numbers[len(numbers)-2]
		numbers = append(numbers, nextNumber)
	}
	sum := 0
	for _, number := range numbers {
		if number%2 == 0 {
			sum += number
		}
	}
	fmt.Println(sum)
}
