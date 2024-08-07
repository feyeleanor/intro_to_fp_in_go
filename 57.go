package main

import (
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a.Add(x)
	}
	os.Exit(int(a))
}

var a Accumulator

type Accumulator int

func (a Accumulator) Add(y int) {
	// This never updates variable a
	a += Accumulator(y)
}
