# Coderbyte Medium Challenges

## 1. Caesar Cipher
Perform a Caesar Cipher shift on it using the num parameter as the shifting number. A Caesar Cipher works by shifting 
each letter in the string N places down in the alphabet (in this case N will be num). Punctuation, spaces, and 
capitalization should remain intact.

For example:

    if the string is "Caesar Cipher" and num is 2 the output should be "Ecguct Ekrjgt"
    if the string is "abc!@# z" and num is 1 the output should be "bcd!@# a"
    if the string is "X^& yZ" and num is 1 the output should be "Y^& zA"

## 2. Formatted Number
Validate a string for a properly formatted number with commas a and return "true" if it is valid and "false".

For example:

    if the string is "12,345.02" the output should be  "true"
    if the string is "12,345.02.1" the output should be  "false"
    if the string is "12345.02" the output should be  "false"

## 3. Number Encoding
Encode every letter into its corresponding numbered position in the alphabet. Punctuation and spaces should remain 
intact.

For example: 
    
    if str is "af5c a#!" then your program should return "1653 1#!".
    if str is "xyz123!@#abc" then your program should return "242526123!@#123".
    if str is "XYZ123 !@#ABC" then your program should return "242526123 !@#123".

## 4. Number Formatting
Convert a string to a properly formatted number.

For example:

    If the string is "12345.02", the output should be "12,345.02"
    If the string is "1212345.02", the output should be "1,212,345.02"
    If the string is "45.02", the output should be "45.02"

## 5. Python Age Counting
Make the http request to https://coderbyte.com/api/challenges/json/age-counting which returns the json data as 
`{"data":"key=IAfpK, age=58, key=WNVdi, age=64}`. Print the total count of age which has a value greater than or equal 
to 50.

For example:

    if the query returns {"data":"key=IAfpK, age=58, key=WNVdi, age=64} then the output should be 2.
    if the query returns {"data":"key=IAfpK, age=40, key=WNVdi, age=50, ey=Unek, age=78} then the output should be 2.
    if the query returns {"data":"key=IAfpK, age=41, key=WNVdi, age=12} then the output should be 0.


## 6. Simple Password
Validate the give password for the following criteria and return "true" if it is valid and "false" otherwise.
 1. There should be at least one Uppercase letter
 2. There should be at least one Number
 3. There should be at least one Punctuation
 4. It should be greater than 7 characters and shorter than 31 characters
 5. Should not contain the word "password"

For example:

    if the string is "passWord!!!", the output should be "false"
    if the string is "adsfA=0g", the output should be "true"
    if the string is "Ukouie&p", the output should be "false"

## 7. Binary Converter
Convert the string of binaries to a decimal number.

For example:

    if the string is "101", the output should be 5
    if the string is "1000", the output should be 8
    if hte string is "1111", the output should be 15
    
## 8. Bracket Matcher / Brace Matcher
Return "true" if the brackets/braces are correctly matched in a string else return "false".

For example:

    if the string is "(hello(there))", the output should be "true"
    if the string is "(hello(there)(world)", the output should be "false"
    if the string is "(hello(t(here)))", the output should be "true"
    
    
 ## 9. Consecutive Numbers
Find the minimum number of integers needed to make the contents of a list consecutive from the lowest number to the 
highest number. Negative numbers may be entered as parameters and no array will have less than 2 elements.
 
 For example:
 
    If array is {4, 8, 6}, the output should be 2. Because, 2 numbers, {5,7} are need between 4 and 8
    If array is {4, 3, 7, 10, 6, 5, 15}, the output should be 6. {8,9,11,12,13,14} are needed
    If array is {21, 28, 20, 25}, the output should be 5. {22,23,24,26,27} are needed
    
## 10. Insert Dashes Asterisks
Insert dashes ('-') between each two odd numbers and insert asterisks ('*') between each two even numbers in a string.
Don't count zero as an odd or even number.
For example: 

    if str is 4546793, the output should be 454*67-9-3
    if str is 23948982934, the output should be 23-94*898*29-34
    if str is 98384763, the output should be 9838*4763
    
## 11. Distinct List (TODO)
In a list of numbers determine the total number of duplicated numbers. If the multiple numbers are duplicated, then 
return the maximum of them.

For example:

    if the input is {1,3,5,5,5,5,7}, the output should be 3. Because there are 3 duplicates of the number 7.
    if the input is {1,2,2,2,3}, the output should be 2. Because there are 2 duplicates of the number 2.
    if the input is {6,9,5,8,9}, the output should be 1. Because there is 1 duplicate of the number 9.

## 12. Greatest Common Factor / Greatest Common Divisor
Find the greatest common divisor or factor of the two numbers and return it.

For example:
 
    if the input is (12,14), the output should be 4
    if the input is (36,6), the output should be 6 
    if the input is (23, 31), the output should be 1 

## 13. Fibonacci Checker
Return the string "yes" if the number given number is a part of the Fibonacci sequence. This sequence is defined by: 
Fn = Fn-1 + Fn-2, which means to find Fn you add the previous two numbers up. The first two numbers are 0 and 1, then 
comes 1, 2, 3, 5 etc. .If num is not in the Fibonacci sequence, return the string "no".

For example:

    if the input is 31, the output should "yes"
    if the input is 58, the output should "no"
    if the input is 144, the output should "yes"

## 14. Formatted Division
Divide num1 by num2, and return the result as a string with properly formatted commas and 4 significant digits after the
decimal place. 

For example:

    if the input is (123456789, 10000), the output should be 12,345.6789
    if the input is (3,4), the output should be 0.7500
    if the input is (100, 33), the output should be 3.3333
 
## 15. Look Say Sequence (TODO: Question not clear)
Find the next number in the sequence according to the following rule: 
Generate the next number in a sequence read off the digits of the given number, counting the number of digits in groups
of the same digit. For example, the sequence beginning with 1 would be: 1, 11, 21, 1211, ...
The 11 comes from there being "one 1" before it and the 21 comes from there being "two 1's" before it. So your program 
should return the next number in the sequence given num.

For example:

    if the input is , the output should be
    if the input is , the output should be
    if the input is , the output should be

## 16. Multiple Brackets
Find the number of brackets if the brackets are correctly matched and each one is accounted for. Otherwise return 0. 
Only "(", ")", "[", and "]" will be used as brackets. If str contains no brackets return 1.
 
For example:
 
    if the input is (hello [world])(!), the output should be 3
    if the input is ((hello [world]), the output should be 0
    if the input is {a(b)c[d]}, the output should be 3

## 17. Number Search
Search for all the numbers in a string, add them together, then return that final number divided by the total number of 
letters in the string rounded to the nearest whole number. Only single digits numbers to be considered.

For example:
 
    if the input is "Hello6 9World 2, Nic8e D7ay!" , the output should be 2 (6 + 9 + 2 + 8 + 7 -> 32/17 ->  1.882 -> 2)
    if the input is "H1i t2h3e4r5e", the output should be 2 ( 1 + 2+ 3 +4 +5 -> 15/7 -> 2.1428 -> 2) 
    if the input is "ho8w ar9e 3y5ou", the output should be 3 ( 8 + 9 + 3 + 5 -> 25/9 -> 2.7777 -> 3)

## 18. Palindrome Two
Check if the string is a Palindrome. Return the string "true" if its a palindrome else "false". Punctuations should be 
ignored.

For example:
 
    if the input is "Anne, I vote more cars race Rome-to-Vienna", the output should be "true"
    if the input is , the output should be
    if the input is , the output should be

## 19. Prime Mover
Return the Nth Prime number between 1 to 10^4. 

For example:
 
    if the input is 1, the output should be 1
    if the input is 5, the output should be 13
    if the input is 10, the output should be 31

## 20. Prime Time
Return the string true if the parameter is a prime number, otherwise return the string false. The range will be 
between 1 and 2^16

For example:
 
    if the input is 144, the output should be "flase"
    if the input is 199, the output should be "true"
    if the input is 89, the output should be "true"

## 21. Simple Mode
Return the number that appears most frequently (the mode).  If there is more than one mode return the one that appeared 
in the array first. If there is no mode return -1. 

For example:
 
    if the input is {10, 4, 5, 2, 4}, the output should be 4
    if the input is {8,4,2,1,9,3}, the output should be -1
    if the input is {1,2,3,4,3,5,6,5,7,3} the output should be 3

## 22. String Scramble
Given two strings return the string true if a portion of str1 characters can be rearranged to match str2, otherwise 
return the string false. Punctuation and symbols will not be entered with the parameters

For example:
 
    if the input is ("rkqodlw", "world") the output should be 2 (wolrd)
    if the input is ("lkasdfeollehasdf", "hello") the output should be "true"
    if the input is ("klajdhfksjdhf", "there"), the output should be

## 23. Swap Two
Given a string, swap the case of each character. Then, if a letter is between two numbers (without separation), then
switch the places of the two numbers

For example:
 
    if the input is "6Hello4 -8World, 7 yes3" , the output should be "4hELLO6 -8wORLD, 7 YES3"
    if the input is "heL1Lo 5theRE Wo7rLd2", the output should be "HEl1Lo 5THEre wO2RlD7"
    if the input is "Hello World", the output should be "hELLO wORLD"

## 24. Three Five Multiples
Given a number return the sum of all the multiples of 3 and 5 that are below the number

For example:
 
    if the input is 10, the output should be 23 (3, 5, 6, 9)
    if the input is 16, the output should be 60 (3,5,6,9,10,12,15)
    if the input is 2, the output should be 0

## 25. Triple Double  
Given two numbers, return 1 if there is a straight triple of a number at any place in num1 and also a straight double 
of the same number in num2. If this isn't the case, return 0.

For example:
 
    if the input is (451999277, 41177722899), the output should be 1
    if the input is (123948555,283745588), the output should be 1
    if the input is (112189374, 41177722899), the output should be 0