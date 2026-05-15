//Write a program that prints the Latin alphabet in lowercase in reverse order (from 'z' to 'a') on a single line.

package main 
import "github.com/01-edu/z01"

func main(){
	
	i := 'z'
	for i>='a'{
		z01.PrintRune(i)
		i--
	} 
	z01.PrintRune('\n')
}