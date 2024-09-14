package sqrt_test

import (
	"fmt"
	"testing"
)

type Ordered interface {
	int | float32 | float64
}

func Sqrt[T Ordered](num T) (float64, error) {
	if num < T(-1) {
		return 0, fmt.Errorf("Can't find a negative value")
	}
	switch any(num).(type) {
	case int:
	case float64:
	case float32:
	default:
		return 0, fmt.Errorf("Unknown type")
	}

	Err := 0.00001
	s := num

	if s == 0 {
		return 0, nil
	}

	for s-num/s > T(Err) {
		s = (s + num/s) / 2
	}

	return float64(s), nil
}

func Abs[T Ordered](num T) (T, error) {
	if num < 0 {
		return num * -1, nil
	}
	return num, nil
}

func almostEqual[T Ordered](v1, v2 T) bool {
	res, err := Abs(v1 - v2)
	if err != nil {
		return false
	}
	return (float64(res) <= 0.001)
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(49)

	if err != nil {
		t.Fatalf("Error in Getting Square Rooot - %f:%v", val, err)
	}

	if !almostEqual(val, 7) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase[T Ordered] struct {
	testvalue     T
	expectedValue T
}

func TestManyFloat(t *testing.T) {
	testCases := []testCase[float64]{
		{0.0, 0.0},
		{4.0, 2.0},
		{2, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.testvalue), func(t *testing.T) {

			val, err := Sqrt(tc.testvalue)
			if err != nil {
				t.Fatalf("Error in Getting Square Rooot - %f:%v", val, err)
			}

			if !almostEqual(val, tc.expectedValue) {
				t.Fatalf("%f != %f", val, tc.expectedValue)
			}
		})
	}
}

func TestManyInt(t *testing.T) {
	testCases := []testCase[int]{
		{0, 0},
		{4, 2},
		{25, 5},
		{9, 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.testvalue), func(t *testing.T) {

			val, err := Sqrt(tc.testvalue)
			if err != nil {
				t.Fatalf("Error in Getting Square Rooot - %f:%v", val, err)
			}

			if !almostEqual(val, float64(tc.expectedValue)) {
				t.Fatalf("%f != %f", val, float64(tc.expectedValue))
			}
		})
	}
}
