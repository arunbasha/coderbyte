//Palindrome Two
//Check if the string is a Palindrome. Return the string "true" if its a palindrome else "false". Punctuations should be
//ignored.
package main

import (
	"fmt"
	"strings"
)

func palindromeTwo(str string) string {
	len := len(str)
	if len < 1 {
		return "false"
	}
	i := 0
	j := len - 1
	str = strings.ToLower(str)
	for j >= i {
		for i <= j && (str[i] < 'a' || str[i] > 'z') {
			i++
		}
		for j >= 0 && (str[j] < 'a' || str[j] > 'z') {
			j--
		}
		if  j == -1 || str[i] != str[j] {
			return "false"
		}
		i++
		j--
	}
	return "true"
}

func  main()  {
	fmt.Println(palindromeTwo("Anne, I vote more cars race Rome-to-Vienna") ) // "true"
	fmt.Println(palindromeTwo("ma! d-a*m")) // "true"
	fmt.Println(palindromeTwo("r~a%ce &ca*r")) // "true"
	fmt.Println(palindromeTwo("helloworld")) // "false"
	fmt.Println(palindromeTwo("howareyou")) // "false"
	fmt.Println(palindromeTwo("hello world there")) // "false"
	fmt.Println(palindromeTwo("rotor")) // "true"
	fmt.Println(palindromeTwo("level")) // "true"
	fmt.Println(palindromeTwo("kayak")) // "true"
	fmt.Println(palindromeTwo("reviver")) // "true"
	fmt.Println(palindromeTwo("!@#$%^&*()(*&^%$#")) // "false"
	fmt.Println(palindromeTwo("z!@#$%^&*()(*&^%$#z")) // "true"
	fmt.Println(palindromeTwo("aa")) // "true"
	fmt.Println(palindromeTwo("a^a")) // "true"
	fmt.Println(palindromeTwo("a@")) // "true"
	fmt.Println(palindromeTwo("a")) // "true"
	fmt.Println(palindromeTwo("*")) // "false"
	fmt.Println(palindromeTwo("")) // "false"
}