//Caesar Cipher
//Perform a Caesar Cipher shift on it using the num parameter as the shifting number. A Caesar Cipher works by shifting
//each letter in the string N places down in the alphabet (in this case N will be num). Punctuation, spaces, and
//capitalization should remain intact.
package main

import (
	"fmt"
	"strings"
)

func caesarCipherV2(input string, key int) string{
	//Cipher main_2.go
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var res []rune

	for _, ch := range input {
		if strings.IndexRune(alphabetUpper, ch) > 0 {
			res = append(res, cipher(ch, []rune(alphabetUpper), key))
		} else if strings.IndexRune(alphabetLower, ch) >= 0 {
			res = append(res, cipher(ch, []rune(alphabetLower), key))
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func cipher(r rune, charset []rune, key int) rune{

	index := strings.IndexRune(string(charset), r)
	index = index + key
	if index >= len(charset) {
		index = index % len(charset)
	}

	// Native approach
	//index := -1
	//for i, ch := range charset {
	//	if ch == r {
	//		index = i
	//	}
	//}
	//
	//if index < 0 {
	//	panic("Index < 0")
	//}
	//
	//for i:=0 ; i< key; i++ {
	//	index++
	//	if index > len(charset) {
	//		index = 0
	//	}
	//}

	return charset[index]
}

func main(){
	fmt.Println(caesarCipherV2("middle-Outz", 2)) // "okffng-Qwvb"
	fmt.Println(caesarCipherV2("Caesar Cipher", 2)) // "Ecguct Ekrjgt"
	fmt.Println(caesarCipherV2("abc!@# z", 1)) // "bcd!@# a"
}