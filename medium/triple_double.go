//Triple Double
//Given two numbers, return 1 if there is a straight triple of a number at any place in num1 and also a straight double
//of the same number in num2. If this isn't the case, return
package main

import "fmt"

func tripleDouble(num1, num2 int) int {
	str1 := fmt.Sprintf("%d", num1)
	str2 := fmt.Sprintf("%d", num2)
	len1 := len(str1)
	len2 := len(str2)
	var num rune
	for i := 0; i < len1 - 2; i++ {
		if str1[i] == str1[i] && str1[i] == str1[i+1] && str1[i] == str1[i+2] {
			num = rune(str1[i])
			break
		}
	}

	for i := 0; i < len2 - 1; i++ {
		if rune(str2[i]) == num && rune(str2[i+1]) == num {
			return 1
		}
	}
	return 0
}

func  main()  {
	fmt.Println(tripleDouble(451999277, 41177722899) )// 1
	fmt.Println(tripleDouble(123948555, 283745588)) // 1
	fmt.Println(tripleDouble(111289374, 41177722899)) // 1
	fmt.Println(tripleDouble(111289374, 412177722899)) // 0
	fmt.Println(tripleDouble(112189374, 41177722899)) // 0
	fmt.Println(tripleDouble(1128983784, 411777822899)) // 0
}