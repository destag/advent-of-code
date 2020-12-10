package year2020

import (
	"advent-of-code/read"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	color string
	num   int
}

func parseBagRule(rule string) (string, []bag) {
	var out []bag

	g := strings.Split(rule, " bags contain ")
	color := g[0]

	re := regexp.MustCompile(`(\d) ([a-z]+ [a-z]+) bags?\.?,? ?`)
	for _, v := range re.FindAllStringSubmatch(rule, -1) {
		n, _ := strconv.Atoi(v[1])
		out = append(out, bag{v[2], n})
	}
	return color, out
}

func findInBags(search string, bags map[string][]bag) []string {
	var contain []string
	for c, b := range bags {
		for _, v := range b {
			if v.color == search {
				fmt.Println(c, " containt ", v.color)
				contain = append(contain, findInBags(c, bags)...)
				fmt.Println(contain)
			}
		}
	}
	return contain
}

func solution07a(data []string) int {
	bags := make(map[string][]bag)
	for _, s := range data {
		c, b := parseBagRule(s)
		bags[c] = b
	}

	f := findInBags("shiny gold", bags)
	fmt.Println(f)
	return -1
}

func solution07b(data []string) int {
	parseBagRule(data[0])

	return -1
}

func solution07(filename string) (int, int) {
	data := read.StringSlice(filename)

	return solution07a(data), solution07b(data)
}
