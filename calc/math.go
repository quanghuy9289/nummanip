package calc

func Add(i int, j int) int {
	return i + j
}

func Add(numbers ... int) int {
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}

	return sum
}