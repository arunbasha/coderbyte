//Multiple Brackets
//Find the number of brackets if the brackets are correctly matched and each one is accounted for. Otherwise return 0.
//Only "(", ")", "[", and "]" will be used as brackets. If str contains no brackets return 1.
package main

import (
	"fmt"
)

func multipleBracketsCount(str string) int {
	var braceCounter, bracketCounter, squareCounter, counter int
	for _, ch := range str {
		switch ch {
		case '{':
			braceCounter++
		case '}':
			if braceCounter > 0 {
				braceCounter--
				counter++
			} else{
					return 0
			}
		case '[':
			squareCounter++
		case ']':
			if squareCounter > 0 {
				squareCounter--
				counter++
			} else {
					return 0
			}
		case '(':
			  bracketCounter++
		case ')':
			if bracketCounter > 0 {
				bracketCounter--
				counter++
			} else {
					return 0
			}
		}
	}

	if bracketCounter != 0 || braceCounter != 0 || squareCounter != 0 {
		return 0
	}
	return counter
}

func  main()  {

	fmt.Println(multipleBracketsCount("(hello [world])(!)")) // 3
	fmt.Println(multipleBracketsCount("((hello [world])")) // 0
	fmt.Println(multipleBracketsCount("{a(b)c[d]}")) // 3
	fmt.Println(multipleBracketsCount("(){}[]{}[]()")) // 6
	fmt.Println(multipleBracketsCount("({[]})"))       // 3
	fmt.Println(multipleBracketsCount("]})[{("))       // 0
	fmt.Println(multipleBracketsCount("({[]}){"))       // 0
	fmt.Println(multipleBracketsCount("({[]})("))       // 0
	fmt.Println(multipleBracketsCount("({{[]})"))       // 0
}