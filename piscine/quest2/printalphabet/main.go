//Write a program that prints the Latin alphabet in lowercase on a single line.
package main

import "github.com/01-edu/z01"


func main(){
	i:='a'
	for i<='z'{
		z01.PrintRune(i)
		i++
	}
	z01.PrintRune('\n')
}