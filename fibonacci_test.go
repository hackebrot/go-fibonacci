package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fibonacciTests = []struct {
	n        int
	expected int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
	{8, 21},
	{9, 34},
	{10, 55},
	{15, 610},
	{20, 6765},
}

func TestIterativeStrategy(t *testing.T) {
	strategy := NewIterative()
	for _, tt := range fibonacciTests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			result := strategy.Compute(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRecursiveStrategy(t *testing.T) {
	strategy := NewRecursive()
	for _, tt := range fibonacciTests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			result := strategy.Compute(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMemoizedStrategy(t *testing.T) {
	strategy := NewMemoized()
	for _, tt := range fibonacciTests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			result := strategy.Compute(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}

	// Test that cache stores computed values
	t.Run("cache_stores_values", func(t *testing.T) {
		memoized := newMemoizedStrategy()
		memoized.Compute(10)
		expectedCache := map[int]int{
			2: 1, 3: 2, 4: 3, 5: 5, 6: 8, 7: 13, 8: 21, 9: 34, 10: 55,
		}
		assert.Equal(t, expectedCache, memoized.memo)
	})
}

func TestConvenienceFunctions(t *testing.T) {
	t.Run("iterative", func(t *testing.T) {
		assert.Equal(t, 55, ComputeIterative(10))
	})
	t.Run("recursive", func(t *testing.T) {
		assert.Equal(t, 55, ComputeRecursive(10))
	})
	t.Run("memoized", func(t *testing.T) {
		assert.Equal(t, 55, ComputeMemoized(10))
	})
}

func TestNegativeInput(t *testing.T) {
	tests := []struct {
		name     string
		strategy Strategy
	}{
		{"iterative", NewIterative()},
		{"recursive", NewRecursive()},
		{"memoized", NewMemoized()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, 0, tt.strategy.Compute(-1))
		})
	}
}

// Benchmark tests - Time complexity: O(n), Space complexity: O(1)
func BenchmarkIterative(b *testing.B) {
	strategy := NewIterative()
	for b.Loop() {
		strategy.Compute(20)
	}
}

// Benchmark tests - Time complexity: O(2^n), Space complexity: O(n)
func BenchmarkRecursive(b *testing.B) {
	strategy := NewRecursive()
	for b.Loop() {
		strategy.Compute(20)
	}
}

// Benchmark tests - Time complexity: O(n), Space complexity: O(n)
func BenchmarkMemoized(b *testing.B) {
	strategy := NewMemoized()
	for b.Loop() {
		strategy.Compute(20)
	}
}
