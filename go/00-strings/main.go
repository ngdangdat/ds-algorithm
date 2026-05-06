package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	nihongo := "日本語"
	fmt.Printf("% 08b\n", []byte(nihongo))
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}
