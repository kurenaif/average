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

func print(w io.Writer, arr []int) {
	if len(arr) == 0 {
		fmt.Fprintf(w, "\n")
		return
	}

	for i, v := range arr {
		fmt.Fprintf(w, "%d", v)
		if i != len(arr)-1 {
			fmt.Fprintf(w, " ")
		} else {
			fmt.Fprintf(w, "\n")
		}
	}
}

func Average(w io.Writer, arr []int) {
	res := make([]int, len(arr))

	for i := 0; i < len(res); i++ {
		res[i] = (arr[i-1] + arr[i] + arr[i+1]) / 3
	}

	print(w, res)
}
