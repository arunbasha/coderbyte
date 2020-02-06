//Three Five Multiples
//Given a number return the sum of all the multiples of 3 and 5 that are below the number
package main

import "fmt"

func threeFiveMultiple(num int) int {
	var sum int
	for n := 1; n < num; n++ {
		if n % 3 == 0 || n %5 == 0 {
			sum += n
		}
	}
	return sum
}

func  main()  {
	fmt.Println(threeFiveMultiple(10) )// 23
	fmt.Println(threeFiveMultiple(16)) // 60
	fmt.Println(threeFiveMultiple(2)) // 0
}