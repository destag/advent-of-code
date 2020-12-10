package util

func Count(data []string, assert func(string) bool) int {
	var sum int
	for _, v := range data {
		if assert(v) {
			sum++
		}
	}
	return sum
}

func Sum(data []string, f func(string) int) int {
	var sum int
	for _, v := range data {
		sum += f(v)
	}
	return sum
}
