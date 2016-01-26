//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func main() {
	counter := 2
	for {
		if divisibleBy(counter) {
			fmt.Println(counter)
			return
		}
		counter++
	}
}

func divisibleBy(number int) bool {
	counter := 2
	for counter <= 20 {
		if number%counter != 0 {
			return false
		}
		counter++
	}
	return true
}
