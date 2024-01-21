package utils_test

import (
	"testing"

	"rpg/utils"
)

// TestSum tests the Sum function in the utils package.
func TestSum(t *testing.T) {
	// Test cases with different scenarios.
	testCases := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Empty slice",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "Negative numbers",
			numbers:  []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
		{
			name:     "Mix of positive and negative numbers",
			numbers:  []int{-1, 2, -3, 4, -5},
			expected: -3,
		},
		{
			name:     "Single number",
			numbers:  []int{42},
			expected: 42,
		},
	}

	// Iterate through each test case and run the test.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Sum function with the current set of numbers.
			result := utils.Sum(tc.numbers)

			// Check if the result matches the expected value.
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d for numbers %v", tc.expected, result, tc.numbers)
			}
		})
	}
}
