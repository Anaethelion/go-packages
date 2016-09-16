package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune:")

	// AppendQuoteRune appends a single-quoted Go character
	// literal representing the rune, as generated by QuoteRune,
	// to dst and returns the extended buffer.
	b = strconv.AppendQuoteRune(b, '☺')

	fmt.Println(string(b))
}