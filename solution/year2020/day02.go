package year2020

import (
	"advent-of-code/read"
	"advent-of-code/util"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	from     int
	to       int
	letter   string
	password string
}

func parseRule(str string) rule {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)
	g := re.FindStringSubmatch(str)

	r := rule{}
	r.from, _ = strconv.Atoi(g[1])
	r.to, _ = strconv.Atoi(g[2])
	r.letter = g[3]
	r.password = g[4]

	return r
}

func solution02a(s string) bool {
	r := parseRule(s)
	c := strings.Count(r.password, r.letter)
	return r.from <= c && c <= r.to
}

func solution02b(s string) bool {
	r := parseRule(s)
	first := string(r.password[r.from-1])
	second := string(r.password[r.to-1])
	return (first == r.letter && second != r.letter) || (first != r.letter && second == r.letter)
}

func solution02(filename string) (int, int) {
	data := read.StringSlice(filename)

	return util.Count(data, solution02a), util.Count(data, solution02b)
}
