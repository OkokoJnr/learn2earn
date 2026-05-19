// Write a function that takes a pointer to an int as argument and gives to this int the value of 1.

package main

import "fmt"



func Pointone(n *int) {
	*n = 1
}
func main(){
	n := 125550
	Pointone(&n)
	fmt.Println(n)
}