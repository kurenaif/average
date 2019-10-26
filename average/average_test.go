package average

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"testing"
)

const EPS = 1e-6

func ApproxEq(lhs []float64, rhs []float64) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for i := 0; i < len(lhs); i++ {
		if math.Abs(lhs[i]-rhs[i]) > EPS {
			return false
		}
	}
	return true
}

func TestAverage(t *testing.T) {
	tests := []struct {
		arr    []int
		expect []float64
	}{
		{[]int{}, []float64{}},
		{[]int{1}, []float64{1.0}},
		{[]int{1, 2}, []float64{1.5, 1.5}},
		{[]int{1, 2, 3}, []float64{1.5, 2.0, 2.5}},
		{[]int{1, 5, 2, 4, 3}, []float64{3.0, 2.666666, 3.6666666, 3.0, 3.5}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("Average(w, %d)", test.arr)
		buf := new(bytes.Buffer)
		Average(buf, test.arr)
		var scanner = bufio.NewScanner(buf)
		scanner.Split(bufio.ScanWords)
		got := []float64{}
		for scanner.Scan() {
			v, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				t.Error(err)
			} else {
				got = append(got, v)
			}
		}

		if !ApproxEq(test.expect, got) {
			t.Errorf("%s = %f, want %f", descr, got, test.expect)
		}
	}
}
