package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// ▁▂▃▄▅▆▇█
func main() {
	path := os.Args[1]
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	str := string(bytes)
	str = regexp.MustCompile("[^GATCgatc]").ReplaceAllString(str, "")
	str = strings.ToUpper(str)
	str = regexp.MustCompile("(.{30})").ReplaceAllString(str, "$1\n")
	str = regexp.MustCompile("G").ReplaceAllString(str, "G▁")
	str = regexp.MustCompile("C").ReplaceAllString(str, "C█")
	str = regexp.MustCompile("A").ReplaceAllString(str, "A▃")
	str = regexp.MustCompile("T").ReplaceAllString(str, "T▆")
	prev := ""
	swapper := regexp.MustCompile("([▁-█])([A-T])")
	for prev != str {
		prev = str
		str = swapper.ReplaceAllString(str, "$2$1")
	}

	str = regexp.MustCompile("([A-T])([▁-█])").ReplaceAllString(str, "$1| $2")
	str = regexp.MustCompile("([A-T▁-█]{5})").ReplaceAllString(str, "$1 ")

	str = regexp.MustCompile(" \n").ReplaceAllString(str, "\n")

	prev = ""
	widener := regexp.MustCompile("\n(.{1,35})\\|")
	for prev != str {
		prev = str
		str = widener.ReplaceAllString(str, "\n$1 |")
	}

	str = regexp.MustCompile("$").ReplaceAllString(str, "\n")
	str = regexp.MustCompile(" \n").ReplaceAllString(str, "\n")
	fmt.Print(str)
}
