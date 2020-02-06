//Return the number that appears most frequently (the mode).  If there is more than one mode return the one that appeared
//in the array first. If there is no mode return -1.
package main

import "fmt"

func simpleMode(num []int) int {
	var prevCounter, currCounter, currNum, prevNum  int
	for i := 0; i< len(num) - 1; i++ {
		currCounter = 0
		for j := i + 1; j < len(num) ; j++ {
			if num[i] == num[j] {
				currCounter++
				currNum = num[i]
			}
		}

		if prevCounter > currCounter {
			currCounter = prevCounter
			currNum = prevNum
		}

		prevNum = currNum
		prevCounter = currCounter
	}

	if currCounter > 0 {
		return currNum
	}
	return -1
}

func  main()  {
	fmt.Println(simpleMode([]int{10, 4, 5, 2, 4}) )// 4
	fmt.Println(simpleMode([]int{8,4,2,1,9,3})) // -1
	fmt.Println(simpleMode([]int{1,2,3,4,3,5,6,5,7,3})) // 3
}