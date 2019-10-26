package main

import (
	"os"

	"./average"
)

func main() {
	average.Average(os.Stdout, []int{})
	average.Average(os.Stdout, []int{1})
	average.Average(os.Stdout, []int{1, 2})
	average.Average(os.Stdout, []int{1, 2, 3})
}
