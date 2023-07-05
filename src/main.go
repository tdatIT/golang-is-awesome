package main

import (
	"LearningGo/src/conditional"
	"LearningGo/src/loop"
	"LearningGo/src/oop"
	"LearningGo/src/variable"
	"fmt"
)

func main() {
	defer println("\nEnd App")
	fmt.Printf("First App Using GoLang\n")
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
	loop.DrawRightTriangle(5)
	fmt.Println("Check Prime Number")
	number := 151
	fmt.Printf("%v is Prime Number: %v", number, loop.IsPrimeNumber(number))
	//Pointer
	/*	number, x := 4, 2
		pointer.Square(number, x)*/
	//Struct
	/*P := struct_go.Fraction{Numerator: 1, Denominator: 2}
	pointer := &P
	pointer.Numerator = 100
	fmt.Println(P)*/
	//Array
	//array.InitArray()
	//OOP
	user := oop.User{Id: 1, Name: "John Nguyen", Age: 24, BMI: 14.24}
	user.ShowInfo()
	user.Run()

	coder := oop.Coder{
		Experiences:         "Senior",
		ProgrammingLanguage: []string{"Java", "C#", "JS", "Golang", "Kotlin"},
		Salary:              1400,
		User:                oop.User{Id: 1, Name: "Dat Tran", Age: 22, BMI: 18.5},
	}
	println("Inheritance\n")
	coder.ShowInfo()
	println("Embedding - Extends\n")
	coder.CreateCV()

}
