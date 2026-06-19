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

func RecursivePower(nb int, power int) int {
	if power == 0{
		return 1
	}
	return nb * RecursivePower(nb, power-1)
}

func main(){
	fmt.Println(IterativePower(4,2))
	fmt.Println(RecursivePower(3,4))
}