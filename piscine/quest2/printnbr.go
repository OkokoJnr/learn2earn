// Write a function that prints an int passed in parameter. All possible values of type int have to go through. You cannot convert to int64.

package main

import "github.com/01-edu/z01"

/*
func PrintNbr(n int){
	if n < 0{
		z01.PrintRune('-')

		if n == -n{
			last := -(n%10)
			PrintNbr(-(n/10))
			z01.PrintRune(rune(last) + '0')
			return
		}


		n = -n
	}
	if n>=10{
		PrintNbr(n/10)
	}

	z01.PrintRune(rune(n%10) + '0')
}
*/
func main() {
	PrintNbr(123)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -n {
			last := -(n % 10)
			n = n / 10

			var digits [20]int
			index := 0
			for n > 0 {
				digits[index] = n % 10
				index++
				n = n / 10
			}
			digits[index] = last
			for j := index - 1; j >= 0; j-- {
				z01.PrintRune(rune(digits[j]) + '0')
			}
		}

		n = -n
	}

	var digits [20]int
	index := 0
	for n > 0 {
		digits[index] = n % 10
		index++
		n = n / 10
	}
	for j := index - 1; j >= 0; j-- {
		z01.PrintRune(rune(digits[j]) + '0')
	}
}
