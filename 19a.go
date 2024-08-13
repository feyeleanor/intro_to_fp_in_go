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
		fmt.Printf("%v: %v\n", i, v)
	}
}
