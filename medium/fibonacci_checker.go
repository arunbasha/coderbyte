//Fibonacci Checker
//Return the string "yes" if the number given number is a part of the Fibonacci sequence. This sequence is defined by:
//Fn = Fn-1 + Fn-2, which means to find Fn you add the previous two numbers up. The first two numbers are 0 and 1, then
//comes 1, 2, 3, 5 etc. .If num is not in the Fibonacci sequence, return the string "no".
package main

import "fmt"

func fibunacciChecker(num int) string {
	if num == 0 || num == 1 {
		return "yes"
	}

	prev := 1
	cur := 1
	for i:= 1; i>0; i++{
		tmp := cur
		cur = prev + cur
		if num > prev && num <= cur {
			if num == cur {
				return "yes"
			} else {
				break
			}
		}
		prev = tmp
	}
	return "no"
}

func  main()  {
	fmt.Println(fibunacciChecker(13))// "yes"
	fmt.Println(fibunacciChecker(14)) // "no "
	fmt.Println(fibunacciChecker(21)) // "yes"
	fmt.Println(fibunacciChecker(22)) // "no"
	fmt.Println(fibunacciChecker(23)) // "no"
	fmt.Println(fibunacciChecker(34)) // "yes"
	fmt.Println(fibunacciChecker(40)) // "no"
	fmt.Println(fibunacciChecker(51)) // "no"
	fmt.Println(fibunacciChecker(55)) // "yes"
	fmt.Println(fibunacciChecker(75)) // "no"
	fmt.Println(fibunacciChecker(89)) // "yes"
	fmt.Println(fibunacciChecker(95)) // "no"
	fmt.Println(fibunacciChecker(124)) // "no"
	fmt.Println(fibunacciChecker(144)) // "yes"
}