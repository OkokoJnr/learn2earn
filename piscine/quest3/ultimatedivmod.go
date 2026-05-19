/*
	Create the following function.

Expected function
func UltimateDivMod(a *int, b *int) {

}
UltimateDivMod should divide the dereferenced value of a by the dereferenced value of b.

Store the result of the division in the int which a points to.

Store the remainder of the division in the int which b points to.
*/
package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	x := *a / *b
	y := *a % *b
	
	*a = x
	*b = y
	
	fmt.Println(x)
	fmt.Println(y)
}
func main(){
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
}