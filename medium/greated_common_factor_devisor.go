//Greatest Common Factor / Greatest Common Divisor
//Find the greatest common divisor or factor of the two numbers and return it.
package main

import "fmt"

func gcdGcf(num1, num2 int) int {
	min := num1
	if num2 > num1 {
		min = num2
	}

	for i:= min; i > 0; i-- {
		if num1 % i ==0 && num2 % i == 0 {
			return i
		}
	}
	return -1
}

func  main()  {
	fmt.Println(gcdGcf(12,16))// 4
	fmt.Println(gcdGcf(12,24)) // 12
	fmt.Println(gcdGcf(36,6)) // 6
	fmt.Println(gcdGcf(6,5)) // 1
}