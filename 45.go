package main
import "os"
import "strconv"

func main() {
	a := MakeAccumulator()
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		a.Add(MakeAccumulator()(x))
	}
	os.Exit(a.Int())
}

type Accumulator func(int) int

func MakeAccumulator() Accumulator {
	var y int
	return func(x int) int {
		y += x
		return y
	}
}

func (a Accumulator) Int() int {
	return a(0)
}

func (a Accumulator) Add(x interface{}) {
	switch x := x.(type) {
	case int:
		a(x)
	case Accumulator:
		a(x.Value())
	}
}