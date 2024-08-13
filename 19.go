package main

import "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	for i, v := range func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	} {
		fmt.Printf("%v: %v\n", i, v)
	}
}
