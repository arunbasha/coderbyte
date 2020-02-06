//Swap Two
//Given a string, swap the case of each character. Then, if a letter is between two numbers (without separation), then
//switch the places of the two numbers
package main

import (
	"fmt"
	"strings"
)

func swapTwo(str string) string {
	var res string
	var numCount, numIndex int
	for index, ch := range str {
		isNewWord := true
		if ch == ' '{
			res = res + string(ch)
			isNewWord = false
			numCount = 0
		} else if ch >= 'A' && ch <= 'Z' {
			res = res + strings.ToLower(string(ch))
		} else if ch >= 'a' && ch <= 'z' {
			res = res + strings.ToUpper(string(ch))
		} else if ch >= '0' && ch <= '9' {
			numCount++
			if numCount == 1 {
				numIndex = index
			}
			if numCount == 2  && isNewWord{
				res = res + string(res[numIndex])
				res = swap(res, numIndex, ch)
			} else {
				res = res + string(ch)
			}
		} else {
			res = res + string(ch)
		}

	}
	return res
}

func swap(str  string, index int, ch rune) string {
	r := []rune(str)
	r[index] = ch
	return string(r)
}


func  main()  {
	fmt.Println(swapTwo("6Hello4 -8World, 7 yes3") )// "4hELLO6 -8wORLD, 7 YES3"
	fmt.Println(swapTwo("heL1Lo 5theRE Wo7rLd2")) // "HEl1Lo 5THEre wO2RlD7"
	fmt.Println(swapTwo("Hello World")) // "hELLO wORLD"
	fmt.Println(swapTwo("H")) // "h"
	fmt.Println(swapTwo("4")) // "4"
	fmt.Println(swapTwo("12")) // "21"
	fmt.Println(swapTwo("1a")) // "1A"
	fmt.Println(swapTwo("1a2 3b4 5c6")) // "2A1 4B3 6C5"
}