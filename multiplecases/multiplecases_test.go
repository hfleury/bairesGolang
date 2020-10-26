package multiplecases

import (
	"fmt"
	"testing"
)

func TestIntMinMultiple(t *testing.T) {
	var cases = []struct {
		a, b int
		want int
	}{
		{3, 4, 3},
		{-9, 5, -9},
		{6, 2, 2},
		{8, 1, 1},
		{10, 20, 20},
	}

	for _, mc := range cases {
		testname := fmt.Sprintf("%d, %d", mc.a, mc.b)
		t.Run(testname, func(t *testing.T) {
			result := IntMin(mc.a, mc.b)
			if result != mc.want {
				t.Errorf("The want %d is not the min value, it is %d : ", mc.want, result)
			}
		})
	}
}
