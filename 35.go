package main
import "os"
import "strconv"

func main() {
	var sum int
	for _, v := range os.Args[1:] {
		x, _ := strconv.Atoi(v)
		sum = add(sum, x)
	}
	os.Exit(sum)
}

func add(x, y int) int {
	return x + y
}