package op_mat

func Add(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Times(a, b int) int {
	return a * b
}

func DividedBy(a, b int) int {
	return a / b
}

func TotalAddition(numbers ...int) int { // <- it applies when you have no idea how many numbers are coming in
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
