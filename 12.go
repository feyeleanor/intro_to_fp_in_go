package main
import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_slice(s, element)
}

func element(s []int, i int) int {
	return s[i]
}

func print_slice(s []int, f func([]int, int) int) {
	defer func() {
		recover()
	}()
	for i := 0; ; i++ {
		Printf("%v: %v\n", i, f(s, i))
	}
}