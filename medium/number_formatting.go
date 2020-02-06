//Number Formatting
//Convert a string to a properly formatted number.
package main

import "fmt"

func numberFormattng(num string) string{

	dotIndex := len(num)
	for i, n := range num {
		if n == '.' {
			dotIndex = i
			break
		}
	}
	if dotIndex < 3 {
		return num
	}

	var res string
	index := 1
	for i := dotIndex - 1; i >= 0; i-- {
		res = string(num[i]) + res
		if  i != 0 && index % 3 == 0 {
			res =  "," + res
		}
		index++
	}

	for i:= dotIndex; i < len(num); i++ {
		res = res + string(num[i])
	}

	return res
}

func main() {
	fmt.Println(numberFormattng("12345.02")) // "12,345.02"
	fmt.Println(numberFormattng("1212345.02")) // "1,212,345.02"
	fmt.Println(numberFormattng("45.02")) // // 45.02
}