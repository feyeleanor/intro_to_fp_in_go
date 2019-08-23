package main
import "os"
import "strconv"

func main() {
	a := MakeAccumulator()
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a(x)
	}
	os.Exit(a(0))
}

func MakeAccumulator() func(int) int {
	var y int
	return func(x int) int {
		y += x
		return y
	}
}