package main
import . "fmt"

func main() {
s := IterableRange{ 0, 2, 5 }
	print_values(s)
	print_values(s)
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

type IterableRange struct {
	start int
	step int
	limit int
}

func (r IterableRange) Each(f func(interface{})) {
	for; r.limit > 0; r.limit-- {
		f(r.start)
		r.start += r.step
	}
}