//Formatted Division
//Divide num1 by num2, and return the result as a string with properly formatted commas and 4 significant digits after the
//decimal place.
package main

import (
	"fmt"
	"strings"
)

func formattedDivision(num1, num2 float64) string {
	var div float64
	div = num1/num2
	str := fmt.Sprintf("%.4f", div)

	dotIndex := strings.Index(str, ".")

	index := 1
	var res string
	for i := dotIndex - 1; i >= 0; i-- {
		if i != 0 && index % 3 == 0 {
			res =  "," + string(str[i]) + res
		} else {
			res = string(str[i]) + res
		}
		index++
	}
	return res + str[dotIndex:]
}

func  main()  {
	fmt.Println(formattedDivision(123456789, 10000))// 12,345.6789
	fmt.Println(formattedDivision(3, 4))// 0.7500
	fmt.Println(formattedDivision(100, 33))// 3.0303
	fmt.Println(formattedDivision(12344,2))// 6,172.0000
	fmt.Println(formattedDivision(23412341,900)) // 	26,013.7122
	fmt.Println(formattedDivision(4314343200,89)) // 	48,475,766.2921
	fmt.Println(formattedDivision(100000000,3)) // 	33,333,333.3333
	fmt.Println(formattedDivision(100,5)) // 	20.0000
	fmt.Println(formattedDivision(10,2)) // 	5.0000
	fmt.Println(formattedDivision(1000,10)) // 	100.0000
	fmt.Println(formattedDivision(0,10)) // 	0.0000
}

