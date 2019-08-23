package main
import . "fmt"

func main() {
	defer func() {
		recover()
	}()
	s := []int{0, 2, 4, 6, 8}
	i := 0
	for {
		Printf("%v: %v\n", i, s[i])
		i++
	}
}