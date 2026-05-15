// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').

package main

import "github.com/01-edu/z01"

func main(){
	i := '0'
for i <='9'{
	z01.PrintRune(i)
	i++
}

z01.PrintRune('\n')
}