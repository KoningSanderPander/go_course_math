package calc

func Add(args ...int) (s int) {
	for _, v := range args {
		s += v
	}
	return
}

func Subtract(x, y int) int {
	return x - y
}
