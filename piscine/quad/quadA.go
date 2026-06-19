/*
Write a function QuadA that prints a valid rectangle with a given width of x and height of y.

The function must draw the rectangles as in the examples.

If x and y are positive numbers, the program should print the rectangles as seen in the examples, otherwise, the function should print nothing.

Make sure you submit all the necessary files to run the program.
*/

package main

import "github.com/01-edu/z01"

func main(){
	QuadA(5,3)
}

func QuadA(x, y int){
for row:=0; row<y; row++{
	for col:=0; col<x; col++{
		first_last_row := row == 0 || row ==y-1
		first_last_col := col == 0 || col ==x-1

		if first_last_row && first_last_col{
			z01.PrintRune('o')
		}else if(first_last_col){
			z01.PrintRune('|')
		}else if (first_last_row){
			z01.PrintRune('-')
		}else if !first_last_row{
			z01.PrintRune(' ')
		}
		}
		z01.PrintRune('\n')
}
}