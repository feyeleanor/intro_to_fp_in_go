package main
import . "reflect"
import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
		print_values(s)
		print_values(func(i int) int {
		return s[i]
	})
}

func print_values(s interface{}) {
	switch s := ValueOf(s); s.Kind() {
	case Func:
		p := make([]Value, 1)
		for_each(func(i int) {
			p[0] = ValueOf(i)
			Printf("%v: %v\n", i, s.Call(p)[0].Interface())
		})
	case Slice:
		for_each(func(i int) {
			Printf("%v: %v\n", i, s.Index(i).Interface())
		})
	}
}

func for_each(f func(int)) (i int) {
	defer func() {
		recover()
	}
	for ; ; i++ {
		f(i)
	}
}