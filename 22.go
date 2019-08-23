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
	defer func() {
		recover()
	}
	switch s := ValueOf(s); s.Kind() {
	case Func:
		for i := 0; ; i++ {
			p := []Value{ ValueOf(i) }
			Printf("%v: %v\n", i, s.Call(p)[0].Interface())
		}
	case Slice:
		for i := 0; ; i++ {
			Printf("%v: %v\n", i, s.Index(i).Interface())
		}
	}
}