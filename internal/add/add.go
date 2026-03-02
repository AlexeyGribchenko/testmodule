package add

func Add(numbers ...int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}
