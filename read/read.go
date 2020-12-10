package read

import (
	"io/ioutil"
	"log"
	"strings"
)

// String return content of file as a string.
func String(filename string) string {
	data := readFile(filename)
	return strings.TrimSpace(string(data))
}

// StringSlice return content of file as a slice of strings.
func StringSlice(filename string) []string {
	data := readFile(filename)
	strs := strings.TrimSpace(string(data))
	return strings.Split(strs, "\n")
}

// Chunks return content of file as a chunks separated by sep.
func Chunks(filename, sep string) []string {
	var data []string
	var cache string
	for _, line := range strings.Split(string(readFile(filename)), "\n") {
		if line == "" {
			data = append(data, strings.TrimSpace(cache))
			cache = ""
			continue
		}
		cache += sep + line
	}
	return data
}

// Decode run provided function on every input line.
func Decode(filename string, call func(string)) {
	data := StringSlice(filename)
	for _, v := range data {
		call(v)
	}
}

func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
