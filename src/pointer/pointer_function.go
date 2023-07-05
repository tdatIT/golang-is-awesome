package pointer

func Square(x int, y int) {
	result := &x
	for i := 0; i < y-1; i++ {
		*result = *result * x
	}
	println(*result)
}
