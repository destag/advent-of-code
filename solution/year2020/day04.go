package year2020

import (
	"advent-of-code/read"
	"advent-of-code/util"
	"regexp"
	"strconv"
	"strings"
)

var required = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

var eyes = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func solution04a(data string) bool {
	m := make(map[string]string)
	s := strings.Split(data, " ")

	for _, v := range s {
		k := strings.Split(v, ":")
		m[k[0]] = k[1]
	}

	for _, v := range required {
		if _, ok := m[v]; !ok {
			return false
		}
	}

	return true
}

func solution04b(data string) bool {
	m := make(map[string]string)
	s := strings.Split(data, " ")

	for _, v := range s {
		k := strings.Split(v, ":")
		m[k[0]] = k[1]
	}

	for _, v := range required {
		if _, ok := m[v]; !ok {
			return false
		}
	}

	byr, _ := strconv.Atoi(m["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	iyr, _ := strconv.Atoi(m["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	eyr, _ := strconv.Atoi(m["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	hgt, _ := m["hgt"]
	hgtu := hgt[len(hgt)-2:]
	hgtv, _ := strconv.Atoi(hgt[:len(hgt)-2])

	if hgtu != "cm" && hgtu != "in" {
		return false
	}

	if hgtu == "cm" && (hgtv < 150 || hgtv > 193) {
		return false
	}

	if hgtu == "in" && (hgtv < 59 || hgtv > 76) {
		return false
	}

	hcl := []byte(m["hcl"])
	if ok, _ := regexp.Match("^#[0-9a-f]{6}$", hcl); !ok {
		return false
	}

	pid := []byte(m["pid"])
	if ok, _ := regexp.Match("^[0-9]{9}$", pid); !ok {
		return false
	}

	ecl := m["ecl"]
	return eyes[ecl]
}

func solution04(filename string) (int, int) {
	data := read.Chunks(filename, " ")

	return util.Count(data, solution04a), util.Count(data, solution04b)
}
