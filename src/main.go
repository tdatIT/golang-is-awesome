package main

import (
	"LearningGo/src/loop"
	"fmt"
)

func main() {
	/*	fmt.Printf("First App Using GoLang\n")
			//Variable
			fmt.Printf("Input radius value:")
			var r int
			_, err := fmt.Scan(&r)
			if err != nil {
				println(err.Error())
			}
			fmt.Printf("Area Circle (%v) is %.2f\n", r, variable.Area(r))
			//Condition
			fmt.Printf("Input month value:")
			var month int
			_, err = fmt.Scan(&month)
			if err != nil {
				println(err.Error())
			}
			fmt.Printf("Month (%v) is (%v)\n", month, conditional.PronounceMonth(month))
		//loop
		fmt.Print("Draw Rectangle\n")
		//loop.DrawRectangle(5, 4)
		loop.DrawRightTriangle(5)*/
	fmt.Println("Check Prime Number")
	number := 151
	fmt.Printf("%v is Prime Number: %v", number, loop.IsPrimeNumber(number))

}
