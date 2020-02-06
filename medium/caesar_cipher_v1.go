//Caesar Cipher
//Perform a Caesar Cipher shift on it using the num parameter as the shifting number. A Caesar Cipher works by shifting
//each letter in the string N places down in the alphabet (in this case N will be num). Punctuation, spaces, and
//capitalization should remain intact.
package main

import "fmt"

func caesarCipherV1(input string, key int) string {
	var res []rune
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			res = append(res,  rotate(ch, 'A', key))
		}else if ch >= 'a' && ch <= 'z' {
			res = append(res, rotate(ch, 'a', key))
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}

func rotate(r rune, base, key int) rune {
	tmp := int(r) - base
	tmp = (tmp + key) % 26
	return rune(tmp + base)
}

//Cipher main.go
func main(){
	fmt.Println(caesarCipherV1("middle-Outz", 2))   // "okffng-Qwvb"
	fmt.Println(caesarCipherV1("Caesar Cipher", 2)) // "Ecguct Ekrjgt"
	fmt.Println(caesarCipherV1("abc!@# z", 1))      // "bcd!@# a"
}
