package array

import "fmt"

func InitArray() {
	n := 0
	_, err := fmt.Scanln(&n)
	if err != nil {
		println(err.Error())
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Input array value Array index (%v): ", i)
		_, err = fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
}
