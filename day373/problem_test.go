package day373

import "testing"

// nolint
var testcases = []struct {
	input []int
	lcs   int
}{
	{[]int{5, 2, 99, 3, 4, 1, 100}, 5},
	{[]int{100, 4, 200, 1, 3, 2}, 4},
	{[]int{100, 4, 101, 103, 1, 102, 3, 104, 2, 105}, 6},
}

func TestLongestConsecutiveSequenceBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		var copied []int

		copy(copied, tc.input)

		if result := LongestConsecutiveSequenceBrute(tc.input); result != tc.lcs {
			t.Errorf("Expected %v, got %v", tc.lcs, result)
		}
	}
}

func BenchmarkLongestConsecutiveSequenceBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			var copied []int

			copy(copied, tc.input)

			LongestConsecutiveSequenceBrute(tc.input)
		}
	}
}
