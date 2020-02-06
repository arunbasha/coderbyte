//Number Search
//Search for all the numbers in a string, add them together, then return that final number divided by the total number of
//letters in the string rounded to the nearest whole number. Only single digits numbers to be considered.
package main

import (
	"fmt"
	"math"
	"strconv"
)

func numberSearch(str string) float64 {
	var numberSum float64
	var letterCount float64
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			n, err := strconv.Atoi(string(ch))
			if err == nil {
				numberSum = numberSum + float64(n)
			}else {
				return 0
			}
		}
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			letterCount++
		}
	}

	var res float64
	res = numberSum / letterCount
	return math.Round(res)
}

func  main()  {
	fmt.Println(numberSearch("Hello6 9World 2, Nic8e D7ay!") )// 2 (6 + 9 + 2 + 8 + 7 -> 32/17 ->  1.8823 -> 2)
	fmt.Println(numberSearch("H1i t2h3e4r5e")) //  2 ( 1 + 2+ 3 +4 +5 -> 15/7 -> 2.1428 -> 2)
	fmt.Println(numberSearch("ho8w ar9e 3y5ou")) // 3 ( 8 + 9 + 3 + 5 -> 25/9 -> 2.7777 -> 3)
}