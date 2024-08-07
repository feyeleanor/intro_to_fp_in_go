package main

import (
	"os"
	"strconv"
)

func main() {
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		accumulate(x)
	}
	os.Exit(y)
}

var y int

func accumulate(x int) {
	y += x
}
