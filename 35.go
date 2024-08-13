package main

import (
	"fmt"
	"math"
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
		switch s.Type().NumIn() {
		case 1: // assume func(int) (T, bool)
			for i := range math.MaxInt {
				p := []reflect.Value{reflect.ValueOf(i)}
				if r := s.Call(p); r[1].Bool() {
					fmt.Printf("1>%v: %v\n", i, s.Call(p)[0].Interface())
				} else {
					return
				}
			}
		case 2: // assume func(func(int, T) bool)
			// TODO:
			//		for i, v := range s {
			//			fmt.Printf("3>%v: %v\n", i, v)
			//		}
		}

	case reflect.Slice:
		for i := range s.Len() {
			fmt.Printf("2>%v: %v\n", i, s.Index(i).Interface())
		}
	}
}
