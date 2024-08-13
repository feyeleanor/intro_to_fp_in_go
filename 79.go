package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var errors int
	computeFactorial := ForValidValues(
		func(i int) {
			fmt.Printf("%v!: %v\n", i, Factorial(i))
		},
		func() {
			if x := recover(); x != nil {
				fmt.Printf("no defined value for %v\n", x)
				errors++
			}
		})
	for _, v := range os.Args[1:] {
		computeFactorial(v)
	}
	os.Exit(errors)
}

func ForValidValues[T Integer](f func(T), e func()) func(string) {
	return func(v string) {
		defer e()

		if x, e := strconv.Atoi(v); e == nil {
			f(T(x))
		} else {
			panic(v)
		}
	}
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
