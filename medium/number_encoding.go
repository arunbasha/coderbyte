//Number Encoding
//Encode every letter into its corresponding numbered position in the alphabet. Punctuation and spaces should remain
//intact.
package main

import (
	"fmt"
)

func numberEncode(input string) string {
	//fmt.Scanf("%s\n",&input)
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//input := scanner.Text()
	var res string

	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			tmp := int(ch) - 'A'
			res = res + fmt.Sprintf("%d", (tmp+1))
		} else if ch >= 'a' && ch <= 'z' {
			tmp := int(ch) - 'a'
			res = res + fmt.Sprintf("%d", (tmp+1))
		} else {
			res = res + string(ch)
		}
	}
	return res

}

func main(){
	fmt.Println(numberEncode("af5c a#!")) // "1653 1#!"
	fmt.Println(numberEncode("xyz123!@#abc")) // "242526123!@#123"
	fmt.Println(numberEncode("XYZ123 !@#ABCz")) // "242526123 !@#12326"
}