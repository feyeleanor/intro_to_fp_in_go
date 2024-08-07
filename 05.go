package main

import . "fmt"

func main() {
	for i, v := range []int{0, 2, 4, 6, 8} {
		Printf("%v: %v\n", i, v)
	}
}
