package main

import . "fmt"

func main() {
	defer func() {
		recover()
	}()
	for i := 0; ; i++ {
		Printf("%v: %v\n", i, []int{0, 2, 4, 6, 8}[i])
	}
}
