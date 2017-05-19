package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var grouping *int
var perLine *int
var filename *string

const (
	LetterG = '▁'
	LetterA = '▃'
	LetterC = '▆'
	LetterT = '█'
)

func Sparkline(r rune) (string, string) {
	switch r {
	case 'G', 'g':
		return string(LetterG), "G"
	case 'A', 'a':
		return string(LetterA), "A"
	case 'C', 'c':
		return string(LetterC), "C"
	case 'T', 't':
		return string(LetterT), "T"
	default:
		return "", ""
	}
}

func printDisplays(display, pic string) {
	displayLen := (*grouping * *perLine) + *perLine
	if display == "" {
		return
	}
	// Pad the display output
	display += strings.Repeat(" ", displayLen-len(display))
	fmt.Println(display + "\t|\t" + pic)
}

func init() {
	grouping = flag.Int("group", 5, "The size of the grouping")
	perLine = flag.Int("line", 10, "The number of the groups per line")
	filename = flag.String("file", "test.seq", "The file to display")
}

func main() {

	flag.Parse()

	f, err := os.Open(*filename)
	defer f.Close()
	if err != nil {
		// Don't do this!!
		panic(err)
	}
	// get
	var buff = make([]byte, 1024*5)
	charCount := 0
	groupCount := 0
	var Display, SparkDisplay string

	for {
		n, err := f.Read(buff)
		if err != nil && err != io.EOF {
			// Don't do this!!
			panic(err)
		}

		if err != nil && err == io.EOF {
			break
		}
		if n == 0 {
			continue
		}
		for _, r := range string(buff[:n]) {
			sparkChar, displayChar := Sparkline(r)
			if sparkChar == "" {
				continue
			}
			Display += displayChar
			SparkDisplay += sparkChar
			charCount++
			if charCount == *grouping {
				Display += " "
				SparkDisplay += " "
				charCount = 0
				groupCount++
			}
			if groupCount == *perLine {
				printDisplays(Display, SparkDisplay)
				Display = ""
				SparkDisplay = ""
				groupCount = 0
			}
		}
	}
	printDisplays(Display, SparkDisplay)
}
