package main
import "os"

func main() {
	os.Exit(add(3, 4))
}

func add(x int, y int) int {
	return x + y
}