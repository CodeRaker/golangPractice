package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	var theArg string
	theArg = os.Args[1]
	fmt.Println("Cool World!", theArg)
	printMoreShit()
}

func printMoreShit() {
	v1, v2, v3 := "1", 2, "3"
	fmt.Printf("%s - %d - %s", v1, v2, v3)
}

// https://gobyexample.com/
// https://play.golang.org/
// https://golang.org/doc/effective_go.html