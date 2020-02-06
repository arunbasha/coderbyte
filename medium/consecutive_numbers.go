//Consecutive Numbers
//Find the minimum number of integers needed to make the contents of a list consecutive from the lowest number to the
//highest number. Negative numbers may be entered as parameters and no array will have less than 2 elements.
package main

import (
	"fmt"
)

func consecutiveNumbers(num []int) int {
	var min, max int
	len := len(num)
	if len > 0 {
		min = num[0]
		max = num[0]
	}

	// Find the min and max numbers
	for _, n := range num {
		if n  < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return max - min - len + 1
}

func  main()  {
	fmt.Println(consecutiveNumbers([]int{4, 8, 6})) // 2
	fmt.Println(consecutiveNumbers([]int{4, 3, 7, 10, 6, 5, 15})) // 6
	fmt.Println(consecutiveNumbers([]int{21, 28, 20, 25})) // 5
}