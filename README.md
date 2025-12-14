# go-fibonacci

Fibonacci strategies in Go

## Installation

```bash
go get github.com/hackebrot/go-fibonacci
```

## Usage

```go
import "github.com/hackebrot/go-fibonacci"

// Convenience functions
fibonacci.ComputeIterative(10)  // 55
fibonacci.ComputeRecursive(10)  // 55
fibonacci.ComputeMemoized(10)   // 55

// Reuse memoized strategy to benefit from cache
strategy := fibonacci.NewMemoized()

strategy.Compute(100)  // computes and caches
strategy.Compute(50)   // reuses cache
```

## Strategies

| Strategy  | Time Complexity | Space Complexity |
|-----------|-----------------|------------------|
| Iterative | O(n)            | O(1)             |
| Recursive | O(2^n)          | O(n)             |
| Memoized  | O(n)            | O(n)             |

## License

MIT
