//Prime Mover
//Return the Nth Prime number between 1 to 10^4.
package main

import "fmt"

func primeMover(num int) int {
	var count int
	max := 10000
	if num == 1 {
		return 2
	}

	for i := 3; i < max; i++ {
		isPrime := true
		for j := 2; j< i; j++ {
			if i % j == 0 {
				isPrime = false
			}
		}

		if isPrime {
			count++
			if num == count {
				return i
			}
		}

	}
	fmt.Println("num: ", num)

	return -1
}

func  main()  {
	fmt.Println(primeMover(1) )// 2
	fmt.Println(primeMover(2) )// 3
	fmt.Println(primeMover(5)) // 13
	fmt.Println(primeMover(10)) // 21
	fmt.Println(primeMover(25)) // 101
	fmt.Println(primeMover(1000)) // 7927
}