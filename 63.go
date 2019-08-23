package main

import "os"
import"strconv"

var limit int

func init() {
	if x, e := strconv.Atoi(os.Args[1]); e == nil {
		limit = x
	}
}

func main() {
	limit--
	if limit > 0 {
		main()
	}
}