package main

import "fmt"

func printStr(s string) {
	fmt.Println(s)
}

func main() {
	s := "default"
	defer func() {
		printStr(s)
	}()
	s = "hello"
}
