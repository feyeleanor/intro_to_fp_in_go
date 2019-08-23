package main
import . "fmt"

func main() {
	print_slice([]int{0, 2, 4, 6, 8})
}

func print_slice(s interface{}) {
	if s, ok := s.([]int); ok {
		for i, v := range s {
			Printf("%v: %v\n", i, v)
		}
	}
}