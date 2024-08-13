package main

import "fmt"

func main() {
	print_slice(0, 2, 4, 6, 8)
}

func print_slice[T any](s ...T) {
	for i, v := range s {
		fmt.Printf("%v: %v\n", i, v)
	}
}
