//Simple Password
//Validate the give password for the following criteria and return "true" if it is valid and "false" otherwise.
//1. There should be at least one Uppercase letter
//2. There should be at least one Number
//3. There should be at least one Punctuation
//4. It should be greater than 7 characters and shorter than 31 characters
//5. Should not contain the word "password"
package main
import "fmt"
import "strings"
import "strconv"

func SimplePassword(str string) string {

	// code goes here
	const punc = `!"#$%&\'()*+, -./:;<=>?@[\]^_{|}~`
	strLen := len(str)
	if strLen <= 7 || strLen >= 31 {
		return "false"
	}

	if strings.Contains(strings.ToLower(str), "password"){
		return "false"
	}

	var containsNum, containsPunc, containsCapital bool
	for _, ch := range str {
		if strings.ToUpper(string(ch)) == string(ch) {
			containsCapital = true
		}
		_, err := strconv.Atoi(string(ch))
		if err != nil {
			containsNum = true
		}
		if strings.Contains(punc, string(ch)) || string(ch) == "`" {
			containsPunc = true
		}
	}
	if containsNum && containsPunc && containsCapital {
		return "true"
	} else {
		return "false"
	}
}

func main() {
	fmt.Println(SimplePassword("turkey90AAA=")) //"true"
	fmt.Println(SimplePassword("passWord123!!!!")) //"false"
	fmt.Println(SimplePassword("passAord123!!!!")) //"true"
	fmt.Println(SimplePassword("password123!!!!")) //"false"
	fmt.Println(SimplePassword("password123!!!!")) //"false"
	fmt.Println(SimplePassword("password!!!!")) //"false"
	fmt.Println(SimplePassword("password")) //"false"
	fmt.Println(SimplePassword("password1!")) //"false"
	fmt.Println(SimplePassword("A1unA1unA1unA1unA1unA1unA1unA!u")) //"false"
	fmt.Println(SimplePassword("A1unA1unA1unA1unA1unA1unA1unA!")) //"true"
}