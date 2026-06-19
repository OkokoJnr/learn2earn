// Instructions
// Write an iterative function that returns the value of nb to the power of power.

// Negative powers will return 0. Overflows do not have to be dealt with.

package main

import "fmt"


func IterativePower(nb int, power int) int {
	res := 1
	for i:=0; i<power; i++{
		res = res * nb
	}
	return res
}


func main(){
	fmt.Println(IterativePower(4,2))
}