//String Scramble
//Given two strings return the string true if a portion of str1 characters can be rearranged to match str2, otherwise
//return the string false. Punctuation and symbols will not be entered with the parameters
package main

import "fmt"

func stringScramble(str1, str2 string) string {
	count :=  make(map[rune]int)
	for _, ch := range str2 {
		var found bool
		var counter int
		count[ch]++
		for _, c := range str1 {
			if ch == c {
				found = true
				counter++
				//fmt.Println(string(ch), " found ", counter, " times", " ch count ", count[ch])
			}
		}
		if !found || count[ch] > counter{
			//fmt.Println("ch ", string(ch), " count[ch] ", count[ch], " counter ", counter)
			return "false"
		}
	}

	return "true"
}

func  main()  {
	fmt.Println(stringScramble("rkqodlw", "world") )// ""
	fmt.Println(stringScramble("lkasdfleolehasdf", "hello")) // ""
	fmt.Println(stringScramble("klajdhfksjdhf", "there")) // ""
	fmt.Println(stringScramble("kola jdeh fkwsjd urhfyo", "how are you")) // ""
	fmt.Println(stringScramble("klajdhfksjdhf", "there")) // ""
	fmt.Println(stringScramble("klajdhfksjdhf", "there")) // ""
}