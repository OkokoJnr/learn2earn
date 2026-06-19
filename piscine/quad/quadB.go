package main

import "github.com/01-edu/z01"

func QuadB(y,x int){
for row:=0; row<x; row++{
	for col:=0; col<y; col++{
		topLeft_bottomRight := (row ==0 && col ==0) || ( row == x-1 && col == y-1)
		topRight_bottomLeft := (row == 0 && col == y-1) || (row == x-1 && col == 0)
		other_rows_cols := (col>0 && col < y-1) && ( row >0 && row < x-1)
		if topLeft_bottomRight {
			z01.PrintRune('/')
		}else if topRight_bottomLeft {
			z01.PrintRune('\\')
		}else if other_rows_cols{
			z01.PrintRune(' ')
		}else{
			z01.PrintRune('*')
		}
				
	}
	z01.PrintRune('\n')
	
}

}

func main(){
	QuadB(5,5)
}
