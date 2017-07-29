package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
}

func manualJoin() {
	s, sep := "", ""
	for _, args := range os.Args[1:] {
		s += sep + args
		sep = " "
	}
	fmt.Println(s)
}

func libraryJoin() {
	fmt.Println(strings.Join(os.Args, " "))
}
