//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

//a2 + b2 = c2
//For example, 32 + 42 = 9 + 16 = 25 = 52.

//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.

package main

import "fmt"

func main() {
	t := possibleTriplets()
	fmt.Println(findTriplet(t))

}

func findTriplet(triplets [][]int) []int {
	var result []int
	for _, triplet := range triplets {
		if triplet[0]*triplet[0]+triplet[1]*triplet[1] == triplet[2]*triplet[2] {
			result = triplet
		}
	}
	return result
}

func possibleTriplets() [][]int {
	var result [][]int
	for a := 1; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := 1000 - a - b
			if b >= c {
				break
			}
			result = append(result, []int{a, b, c})
		}
	}
	return result
}
