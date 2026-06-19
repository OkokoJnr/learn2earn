package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 1{
		return 0
	}

	return nb * IterativeFactorial(nb-1)
}

func main() {
	fmt.Println(IterativeFactorial('k'))
}