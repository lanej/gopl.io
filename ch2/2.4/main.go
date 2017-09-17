// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.
package main

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x using a loop
func PopCount(x uint64) int {
	var count byte

	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			count++
		}
	}

	return int(count)
}

// PopCountLoop returns the population count (number of set bits) of x using a loop
func PopCountLoop(x uint64) int {
	var count byte

	for i, v := uint64(0), uint64(8); i < v; i++ {
		count += byte(x >> (i * 8))
	}
	return int(count)
}

//!-
