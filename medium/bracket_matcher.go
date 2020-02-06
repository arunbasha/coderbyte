//Bracket Matcher / Brace Matcher
//Return "true" if the brackets/braces are correctly matched in a string else return "false".
package main

import "fmt"

func bracketMatcher(str string) string {
	var braceCounter, parenthesisCounter, squareBracketCounter int
	for _, ch := range str {
		switch ch {
		case '(':
			parenthesisCounter++
		case ')':
			if parenthesisCounter > 0 {
				parenthesisCounter--
			} else {
				return "false"
			}
		case '{':
			braceCounter++
		case '}':
			if braceCounter > 0 {
				braceCounter--
			} else {
				return "false"
			}
		case '[':
			squareBracketCounter++
		case ']':
			if squareBracketCounter > 0 {
				squareBracketCounter--
			} else {
				return "false"
			}
		}
	}
	if braceCounter != 0 || parenthesisCounter != 0 || squareBracketCounter != 0 {
	 	return "false"
	}

	return "true"

}

func  main()  {
	fmt.Println(bracketMatcher("(hello(there))")) // "true"
	fmt.Println(bracketMatcher("(hello(t(here)))")) // "true"
	fmt.Println(bracketMatcher("())(()")) // "false"
	fmt.Println(bracketMatcher("())(()")) // "false"
	fmt.Println(bracketMatcher("())(())")) // "false"
	fmt.Println(bracketMatcher("())(())")) // "false"
	fmt.Println(bracketMatcher("))))((((")) // "false"
	fmt.Println(bracketMatcher("))(())((")) // "false"
	fmt.Println(bracketMatcher("(())")) // "false"
	fmt.Println()
	fmt.Println(bracketMatcher("{hello{there}}")) // "true"
	fmt.Println(bracketMatcher("{hello{t{here}}}")) // "true"
	fmt.Println(bracketMatcher("{}}{{}")) // "false"
	fmt.Println(bracketMatcher("{}{{}")) // "false"
	fmt.Println(bracketMatcher("{}}{{}}")) // "false"
	fmt.Println(bracketMatcher("{}}{{}}")) // "false"
	fmt.Println(bracketMatcher("}}}}{{{{")) // "false"
	fmt.Println(bracketMatcher("{}{{}}{{")) // "false"
	fmt.Println(bracketMatcher("{{}}")) // "false"
	fmt.Println()
	fmt.Println(bracketMatcher("[hello[there]]")) // "true"
	fmt.Println(bracketMatcher("[hello[t[here]]]")) // "true"
	fmt.Println(bracketMatcher("[]][[]")) // "false"
	fmt.Println(bracketMatcher("[][[]")) // "false"
	fmt.Println(bracketMatcher("[]][[]]")) // "false"
	fmt.Println(bracketMatcher("[]][[]]")) // "false"
	fmt.Println(bracketMatcher("]]]][[[[")) // "false"
	fmt.Println(bracketMatcher("[][[]][[")) // "false"
	fmt.Println(bracketMatcher("[[]]")) // "false"
}
