package fibonacci

// Strategy defines the interface for different Fibonacci computation strategies
type Strategy interface {
	// Compute calculates the nth Fibonacci number
	Compute(n int) int
}

// iterativeStrategy computes Fibonacci numbers using iteration.
// Time: O(n), Space: O(1)
type iterativeStrategy struct{}

func (s *iterativeStrategy) Compute(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// recursiveStrategy computes Fibonacci numbers using pure recursion.
// Time: O(2^n), Space: O(n)
type recursiveStrategy struct{}

func (s *recursiveStrategy) Compute(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return s.Compute(n-1) + s.Compute(n-2)
}

// memoizedStrategy computes Fibonacci numbers using recursion with caching.
// Cache is maintained across multiple calls to the same instance.
// Time: O(n), Space: O(n)
type memoizedStrategy struct {
	memo map[int]int
}

func newMemoizedStrategy() *memoizedStrategy {
	return &memoizedStrategy{
		memo: make(map[int]int),
	}
}

func (s *memoizedStrategy) Compute(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if val, ok := s.memo[n]; ok {
		return val
	}

	result := s.Compute(n-1) + s.Compute(n-2)
	s.memo[n] = result

	return result
}

// NewIterative creates an iterative Fibonacci strategy.
func NewIterative() Strategy {
	return &iterativeStrategy{}
}

// NewRecursive creates a recursive Fibonacci strategy.
func NewRecursive() Strategy {
	return &recursiveStrategy{}
}

// NewMemoized creates a memoized recursive strategy.
// Reuse the instance to benefit from caching across multiple calls.
func NewMemoized() Strategy {
	return newMemoizedStrategy()
}

// ComputeIterative is a convenience function for iterative computation
func ComputeIterative(n int) int {
	return NewIterative().Compute(n)
}

// ComputeRecursive is a convenience function for recursive computation
func ComputeRecursive(n int) int {
	return NewRecursive().Compute(n)
}

// ComputeMemoized is a convenience function for memoized computation
func ComputeMemoized(n int) int {
	return NewMemoized().Compute(n)
}
