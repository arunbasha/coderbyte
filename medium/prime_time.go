//Return the string true if the parameter is a prime number, otherwise return the string false. The range will be
//between 1 and 2^16
package main

import "fmt"

func isPrime(num int) string {
	for i := 2 ; i< num / 2; i++ {
		if num % i == 0 {
			return "false"
		}
	}
	return "true"
}

func  main() {
	fmt.Println(isPrime(144) )// "false"
	fmt.Println(isPrime(89)) // "true"
	fmt.Println(isPrime(199)) // "true"
	fmt.Println(isPrime(1)) // "true"
	fmt.Println(isPrime(2)) // "true"
	fmt.Println(isPrime(3)) // "true"
	fmt.Println(isPrime(5)) // "true"
	fmt.Println(isPrime(7927)) // "true"
	fmt.Println(isPrime(101)) // "true"
}