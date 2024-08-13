package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values(s)
	print_values(func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
	print_values(func(g func(i, v int) bool) {
		for i, v := range s {
			if !g(i, v) {
				return
			}
		}
	})
}

func print_values(s any) {
	switch s := s.(type) {
	case func(int) (int, bool):
		for i := range math.MaxInt {
			if v, ok := s(i); ok {
				fmt.Printf("1>%v: %v\n", i, v)
			} else {
				return
			}
		}
	case []int:
		for i, v := range s {
			fmt.Printf("2>%v: %v\n", i, v)
		}
	case func(func(int, int) bool):
		for i, v := range s {
			fmt.Printf("3>%v: %v\n", i, v)
		}
	}
}
