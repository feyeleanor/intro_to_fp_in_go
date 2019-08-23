package main
import "os"
import "strconv"

func main() {
	x, _ := strconv.Atoi(os.Args[1])
	y, _ := strconv.Atoi(os.Args[2])
	os.Exit(add(x, y))
}

func add(x, y int) int {
	return x + y
}