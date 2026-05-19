// Write a function that prints one by one the characters of a string on the screen.

package main

import (
	
	"github.com/01-edu/z01"
)

func PrintStr(s string){
	slice := []rune(s)
	for i:=0; i<len(slice); i++{
		z01.PrintRune(slice[i])
	}
}
func main(){
	PrintStr("Hello World")
}