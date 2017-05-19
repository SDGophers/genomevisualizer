package main

import (
	"io/ioutil"
	"fmt"
	"bytes"
)

func byteToBar(b byte)  rune {
	switch b {
	case 'A':
		return '▁' // 1/8 bar
	case 'T' :
		return '▃' // 1/4 bar
	case 'C' :
		return '▅'
	case 'G':
		return '█'
	}

	return '\u0000'
}

func main() {
	dat, err := ioutil.ReadFile("sequence.txt")
	if err != nil {
		panic(err)
	}

	dat = bytes.ToUpper(dat)

	valid := 0
	bars := [30]rune{}
	line := 1

	fmt.Printf( "%v\t |", line)
	for _, ch := range dat {
		if !(ch == 'A' || ch == 'T' || ch == 'G' || ch == 'C') {
			continue
		}

			fmt.Print(string(ch))
			bars[valid] = byteToBar(ch)
			valid += 1
			if valid % 5 == 0 {
				fmt.Print(" ")
			}



		if valid == 30 {
			//print bars
			fmt.Print("\t")
			for i, rn := range bars {
				fmt.Print(string(rn))
				if i % 5 == 4 {
					fmt.Print(" ")
				}
			}
			valid = 0
			line += 1
			fmt.Printf( "\n%v\t |", line)

		}
	}

	if valid != 0 {
		left := (30 - valid) / 4
		for i := 0; i <= left ; i += 1 {
			fmt.Print("\t")
		}
		fmt.Print("\t")
		for i, rn := range bars[:valid] {
			fmt.Print(string(rn))
			if i % 5 == 4 {
				fmt.Print(" ")
			}
		}
		fmt.Print("")
	}

}