package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	numChars := 10

	dat, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	sanitized := sanitize(string(dat))

	b := make([]rune, numChars)
	for i, c := range sanitized {
		b[i%numChars] = c
		if i%numChars == 9 {
			printPretty(b)

		}
	}

	// fmt.Println(sanitized)
}

func printPretty(c []rune) {
	fmt.Printf("Rune is: %s %s %s %s\n", string(c[0:5]), string(c[5:]), string(printPrettyBar(c[0:5])), string(printPrettyBar(c[5:])))
}

func printPrettyBar(c []rune) string {
	m := map[string]string{
		"a": "▁",
		"g": "█",
		"t": "▃",
		"c": "▆",
	}

	var l []string
	for _, r := range c {
		l = append(l, m[string(r)])
	}

	return strings.Join(l, "")
}

func readFile() ([]byte, error) {
	path := "sample.txt"

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return dat, nil
}

func sanitize(s string) string {
	lower := strings.ToLower(s)

	nospaces := strings.Replace(lower, " ", "", -1)

	return nospaces
}
