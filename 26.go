package main
import . "fmt"

func main() {
	var s Iterable = IterableSlice{ 0, 2, 4, 6, 8 }
	i := 0
	s.Each(func(v interface{}) {
		Printf("%v: %v\n", i, v)
		i++
	})
}

type Iterable interface {
	Each(func(interface{}))
}

type IterableSlice	[]int

func (i IterableSlice) Each(f func(interface{})) {
	for _, v := range i {
		f(v)
	}
}
