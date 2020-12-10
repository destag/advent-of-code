package year2020

import (
	"advent-of-code/read"
	"log"
	"strconv"
)

func solution01a(data []int) int {
	for _, i := range data {
		for _, j := range data {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return -1
}

func solution01b(data []int) int {
	for _, i := range data {
		for _, j := range data {
			for _, k := range data {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return -1
}

func solution01(filename string) (int, int) {
	var data []int
	read.Decode(filename, func(s string) {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, n)
	})

	return solution01a(data), solution01b(data)
}
