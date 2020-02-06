//Binary Converter
//Convert the string of binaries to a decimal number.
package main

import "fmt"

func binaryConverterLittleEndian(str string) int {
	var res int
	len  := len(str)
	for i := len - 1; i >= 0; i-- {
		if str[i] == '1' {
			res += power(2, len - 1 - i)
		}
	}
	return res
}

func binaryConverterBigEndian(str string) int {
	var res int
	for i, ch := range str {
		if ch == '1' {
			res += power(2, i)
		}
	}
	return res
}

func power(base, power int)  int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * base
	}
	return res
}

func  main()  {
	fmt.Println(binaryConverterLittleEndian("101")) // 5
	fmt.Println(binaryConverterLittleEndian("1000")) // 8
	fmt.Println(binaryConverterLittleEndian("1111")) // 15
	fmt.Println(binaryConverterLittleEndian("1111111")) // 127

	fmt.Println(binaryConverterBigEndian("101")) // 5
	fmt.Println(binaryConverterBigEndian("1000")) // 8
	fmt.Println(binaryConverterBigEndian("1111")) // 15
	fmt.Println(binaryConverterBigEndian("1111111")) // 127

}