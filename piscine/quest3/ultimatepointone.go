// Write a function that takes a pointer to a pointer to a pointer to an int as argument and gives to this int the value of 1.

package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}


func main(){
	x := 3
	y := &x
	z := &y
	UltimatePointOne(&z)
	fmt.Println(x)
}