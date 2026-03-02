package add

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -15},
		{"mixed numbers", []int{-1, 2, -3, 4, -5}, -3},
		{"with zeros", []int{0, 0, 0}, 0},
		{"single number", []int{42}, 42},
		{"empty arguments", []int{}, 0},
		{"large numbers", []int{1000000, 2000000, 3000000}, 6000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.numbers...)
			if result != tt.expected {
				t.Errorf("Add(%v) = %d; expected %d", tt.numbers, result, tt.expected)
			}
		})
	}
}
