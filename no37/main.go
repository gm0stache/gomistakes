package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hêllo"

	fmt.Println("naive iteration")
	fmt.Printf("string: %s, len: %d\n", s, len(s))
	for i := range s {
		fmt.Printf("idx: %d, value: %c\n", i, s[i])
	}

	fmt.Println("")
	fmt.Println("rune aware iteration")
	fmt.Printf("string: %s, len: %d\n", s, utf8.RuneCountInString(s))
	for i, v := range s {
		fmt.Printf("idx: %d, value: %c\n", i, v)
	}
}
