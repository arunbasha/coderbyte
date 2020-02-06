//Insert Dashes Asterisks
//Insert dashes ('-') between each two odd numbers and insert asterisks ('*') between each two even numbers in a string.
//Don't count zero as an odd or even number.
package main

import "fmt"

func insertDashesAsterisks(str string) string {
	var res string
	len := len(str)
	for i := 0; i< len - 1; i++ {
		if str[i] % 2 == 0 &&  str[i+1] % 2 == 0{
			res = res + string(str[i]) + "*"
		} else if str[i] % 2 != 0 &&  str[i+1] % 2 != 0{
			res = res + string(str[i]) + "-"
		} else {
			res = res + string(str[i])
		}
	}
	res = res + string(str[len-1])
	return res
}

func  main()  {
	fmt.Println(insertDashesAsterisks("4546793")) // "454*67-9-3"
	fmt.Println(insertDashesAsterisks("23948982934")) // "23-94*898*29-34"
	fmt.Println(insertDashesAsterisks("98384763")) // "9838*4763"
}