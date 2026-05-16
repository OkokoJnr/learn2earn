// Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.

// These combinations are separated by a comma and a space.


package main

 import "github.com/01-edu/z01"


func main(){
	PrintComb2()
}

func PrintComb2(){
	for i:=0; i<=99; i++{
		for j:=i+1; j<=99; j++{
			if (i != j){
				t1,t2:= rune(i/10), rune(j/10)
				r1,r2 := rune(i%10), rune(j%10)
				z01.PrintRune(t1 + '0')
				z01.PrintRune(r1 + '0')
				z01.PrintRune(' ')
				z01.PrintRune(t2 + '0')
				z01.PrintRune(r2 + '0')
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
