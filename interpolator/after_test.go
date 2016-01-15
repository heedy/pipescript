package interpolator

import "testing"

func TestAfter(t *testing.T) {
	TestCase{
		Interpolator: "after",
		Input:        testDpa,
		Output: []TestOutput{
			{0.5, &testDpa[0]},
			{0.7, &testDpa[0]},
			{2.0, &testDpa[2]},
			{5.5, &testDpa[5]},
			{6.0, &testDpa[7]},
			{8.0, nil},
		},
	}.Run(t)
}
