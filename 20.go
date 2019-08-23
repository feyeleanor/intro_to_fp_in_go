package main
import "reflect"
import . "fmt"

func main() {
	s := []int{0, 2, 4, 6, 8}
		print_values(s)
		print_values(func(i int) int {
		return s[i]
	})
}

func print_values(s interface{}) {
	switch s := reflect.ValueOf(s); s.Kind() {
	case reflect.Func:
		for i := 0; i < 5; i++ {
			p := []reflect.Value{ reflect.ValueOf(i) }
			Printf("%v: %v\n", i, s.Call(p)[0].Interface())
		}
	case reflect.Slice:
		for i := 0; i < s.Len(); i++ {
			Printf("%v: %v\n", i, s.Index(i).Interface())
		}
	}
}