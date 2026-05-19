// Write a function that takes two pointers to an int (*int) and swaps their contents.

package main

import "fmt"

func Swap(a *int, b *int){
	*a, *b = *b, *a
}

func main(){
	a := 0
	b := 1
	fmt.Println(a)
	fmt.Println(b)
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}