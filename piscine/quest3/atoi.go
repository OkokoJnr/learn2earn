/*
Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

For this exercise the handling of the signs + or - does have to be taken into account.

This function will only have to return the int. For this exercise the error result of Atoi is not required.
*/
package main
import "fmt"

func Atoi(s string) int{
	res := 0
	sign := 1
	for i, ch := range s{
		if i == 0 {
			if ch == '-'{
				sign = -1
				continue
			}
			if ch == '+'{
				sign = +1
				continue
			}
			if ch<'0' || ch>'9'{
				return 0
			}
		}else{
			if ch<'0' || ch>'9'{
				return 0
			}
		}
		res = res * 10 + int(ch-'0')

		}
		return res * sign
}

func main(){
	fmt.Println(Atoi("k0000000012345"))
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}