package main

import (
	"fmt"
	"strings"
)

func main() {
	counts:=map[string]int{}
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words:=strings.Fields(text)
	for _, value := range words {
		counts[strings.ToLower(value)]++
	}

	fmt.Println(counts)
}
