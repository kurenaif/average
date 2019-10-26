package average

import (
	"fmt"
	"io"
)

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func max(lhs int, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func print(w io.Writer, arr []float64) {
	for i, v := range arr {
		fmt.Fprintf(w, "%f", v)
		if i != len(arr)-1 {
			fmt.Fprintf(w, " ")
		} else {
			fmt.Fprintf(w, "\n")
		}
	}
}

func Average(w io.Writer, arr []int) {
	res := make([]float64, len(arr))

	if len(arr) == 0 {
		print(w, res)
		return
	}
	if len(arr) == 1 {
		res[0] = float64(arr[0])
		print(w, res)
		return
	}

	res[0] = float64(arr[0]+arr[1]) / 2
	for i := 1; i < len(res)-1; i++ {
		res[i] = float64(arr[i-1]+arr[i]+arr[i+1]) / 3
	}
	res[len(arr)-1] = float64(arr[len(arr)-1]+arr[len(arr)-2]) / 2

	print(w, res)
}
