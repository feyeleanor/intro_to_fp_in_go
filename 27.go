package main
import . "fmt"

func main() {
	print_values(IterableLimit(5))
	print_values(IterableSlice{ 0, 2, 4, 6, 8 })
}

func print_values(s Iterable) (i int) {
	s.Each(func(v interface{}) {
		Printf("%v: %v\n", i, v)
		i++
	})
	return i
}

type Iterable interface {
	Each(func(interface{}))
}

type IterableLimit	int
type IterableSlice	[]int

func (i IterableLimit) Each(f func(interface{})) {
	for v := IterableLimit(0); v < i; v++ {
		f(v)
	}
}

func (i IterableSlice) Each(f func(interface{})) {
	for _, v := range i {
		f(v)
	}
}