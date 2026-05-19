/*
Write a function that reverses a string.

This function will return the reversed string.
*/

package main

import "fmt"

func StrRev(s string) string{
	slice_s := []byte(s)
	revstr := ""

	for i:=len(slice_s)-1; i>=0; i--{
		revstr += string(slice_s[i])
	}
	return revstr
}



func StrRev1(s string) string{
	r := []rune(s)
	for i,j := 0, len(r)-1;i<j; i,j = i+1, j-1{
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func main(){
	fmt.Println(StrRev("😊Hello World! é"))
	fmt.Println(StrRev1("😊Hello World! é"))
}