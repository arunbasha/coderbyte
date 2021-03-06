//Formatted Number
////Validate a string for a properly formatted number with commas a and return "true" if it is valid and "false".
package main

import (
	"fmt"
	"strings"
)

func FormattedNumberV2(strArr []string) string {
	// code goes here
	res := "true"
	for _, val := range strArr {
		dotIndex := strings.LastIndex(val, ".")
		if strings.Index(val,".") != dotIndex {
			return "false"
		}
		str := []rune(val)
		for i := dotIndex + 1; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				return "false"
			}
		}

		index := 1
		commaFound := false
		for i := dotIndex - 1; i >= 0; i-- {
			if str[i] < '0' || str[i] > '9' {
				if string(str[i]) != "," {
					return "false"
				}
			}
			if i != 0  && str[i] == ',' {
					commaFound = true
					if index%4 != 0 {
						return "false"
					}
			}
			index++
		}
		if index > 4 && !commaFound {
			return "false"
		}
	}
	return res
}

func main() {
	fmt.Println(FormattedNumberV2([]string {"2,561.002"}))    // "true"
	fmt.Println(FormattedNumberV2([]string {"2,56a.002"}))    // "false"
	fmt.Println(FormattedNumberV2([]string {"2,560.00.2"}))   // "false"
	fmt.Println(FormattedNumberV2([]string {"0.232567"}))     // "true"
	fmt.Println(FormattedNumberV2([]string {"1,093,222.04"})) // "true"
	fmt.Println(FormattedNumberV2([]string {"109,322.04"}))   // "true"
	fmt.Println(FormattedNumberV2([]string {"109,322.04a"}))  // "false"
	fmt.Println(FormattedNumberV2([]string {"1,093,22.04"}))  // "false"
	fmt.Println(FormattedNumberV2([]string {"10,932,2.04"}))  // "false"
	fmt.Println(FormattedNumberV2([]string {"10,9,322.04"}))  // "false"
}