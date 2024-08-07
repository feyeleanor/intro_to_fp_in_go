package main

import (
	. "fmt"
	"reflect"
)

func main() {
	s := []int{0, 2, 4, 6, 8}
	print_values(s)
	print_values(func(i int) (v int, ok bool) {
		if ok = (-1 < i) && (i < len(s)); ok {
			v = s[i]
		}
		return
	})
}

func print_values(s any) {
	switch s := reflect.ValueOf(s); s.Kind() {
	case reflect.Func:
		for i := 0; ; i++ {
			p := []reflect.Value{reflect.ValueOf(i)}
			if r := s.Call(p); r[1].Bool() {
				Printf("5>s(%v): %v\n", i, s.Call(p)[0].Interface())
			} else {
				return
			}
		}
	case reflect.Slice:
		for i := 0; i < s.Len(); i++ {
			Printf("6>s[%v]: %v\n", i, s.Index(i).Interface())
		}
	}
}
