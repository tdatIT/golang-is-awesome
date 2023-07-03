package loop

import (
	"fmt"
)

func DrawRectangle(a int, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}

func DrawRightTriangle(x int) {
	for i := 0; i < x; i++ {
		for j := 0; j <= i; j++ {
			print("*")
		}
		print("\n")
	}
}
func IsPrimeNumber(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return x > 1
}
