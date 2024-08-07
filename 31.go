package main

import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values(s)
	print_values(func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
}

func print_values(s any) {
	switch s := s.(type) {
	case func(int) (int, bool):
		for i := 0; ; i++ {
			if v, ok := s(i); ok {
				Printf("s(%v): %v\n", i, v)
			} else {
				return
			}
		}
	case []int:
		for i, v := range s {
			Printf("s[%v]: %v\n", i, v)
		}
	}
}
