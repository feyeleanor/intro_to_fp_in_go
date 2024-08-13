package main

import "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	f := func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	}
	for i, v := range f {
		fmt.Printf("1>%v: %v\n", i, v)
	}

	for i, v := range each(s) {
		fmt.Printf("2>%v: %v\n", i, v)
	}
}

func each(s []int) func(func(int, int) bool) {
	return func(g func(int, int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	}
}
