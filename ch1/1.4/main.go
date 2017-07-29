package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[*os.File]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, countmap := range counts {
		if len(countmap) > 1 {
			fmt.Printf("%q\t", line)
			for file := range countmap {
				fmt.Printf(" %s", file.Name())
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]map[*os.File]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[input.Text()]; !ok {
			counts[input.Text()] = make(map[*os.File]int)
		}
		counts[input.Text()][f]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
