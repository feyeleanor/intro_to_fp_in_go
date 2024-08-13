package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("no factorial")
		}
	}()
	x, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("%v!: %v\n", x, Factorial(x))
}

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func Factorial[T Integer](n T) (r T) {
	if n < 0 {
		panic(n)
	}
	for r = 1; n > 0; n-- {
		r *= n
	}
	return
}
