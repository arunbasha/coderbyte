package main

import "fmt"

func duplicatedNumbers(num []int) int {
	var maxCount int
	len := len(num)
	for i := 0; i < len-1; i++ {
		var curMaxCount int
		for j := i; j < len; j++ {
			if num[i] == num[j] {
				curMaxCount++
			}
		}
		if curMaxCount > maxCount {
			maxCount = curMaxCount
		}
	}
	return maxCount - 1
}

func  main()  {
	fmt.Println(duplicatedNumbers([]int{1,3,5,5,5,5,7})) // 3
	fmt.Println(duplicatedNumbers([]int{1,2,2,2,3})) // 2
	fmt.Println(duplicatedNumbers([]int{6,9,5,8,9})) // 1
	fmt.Println(duplicatedNumbers([]int{6,9,5,8,9,6,6})) // 2
	fmt.Println(duplicatedNumbers([]int{8,6,8,9,5,8,9,8, 6,8,6,8})) // 5
}