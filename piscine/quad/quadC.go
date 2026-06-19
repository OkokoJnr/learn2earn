package main

import "github.com/01-edu/z01"

func QuadC(x, y int){
	for row:=0; row<x; row++{
		for col:=0; col< y; col++{
			topLeftRight := row == 0 && (col == 0 || col == y-1)
			bottomLeftRight := row == x-1 && (col == 0 || col == y-1)
			others := (row > 0 && row < x-1) && (col >0 && col < y-1)
			if topLeftRight {
				z01.PrintRune('A')
			}else if bottomLeftRight{
				z01.PrintRune('C')
			}else if others{
				z01.PrintRune(' ')
			}else{
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
	}
}
func main(){
	QuadC(1,5)
}