// Write a function that counts the runes of a string and that returns that count

package main

import "fmt"


func StrLen(s string) int {
	slice_s := []rune(s)
	count :=0
	for i:=0; i<len(slice_s); i++{
		count ++
	}
	return count
}

func main(){
	fmt.Println(StrLen("Hello World! "))
}