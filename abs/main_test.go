package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{
			name:     "Integer",
			value:    3,
			expected: 3,
		},
		{
			name:     "Negative integer",
			value:    -3,
			expected: 3,
		},
		{
			name:     "Negative float",
			value:    -2.000001,
			expected: 2.000001,
		},
		{
			name:     "Small float",
			value:    -0.000000003,
			expected: 0.000000003,
		},
		{name: "Negative zero", value: -0, expected: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := Abs(test.value); result != test.expected {
				t.Errorf("For value %v expected result %v not equal with actual %v", test.value, test.expected, result)
			}
		})
	}
}
