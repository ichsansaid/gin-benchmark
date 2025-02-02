package ucases

import (
	"context"

	"github.com/ichsansaid/golang-gin-benchmark/internal/contracts"
)

// FiboRecursiveCPUIntensiveUcase implements FiboRecursiveCPUIntensiveUcaseContract
type FiboRecursiveCPUIntensiveUcase struct{}

// NewFiboRecursiveCPUIntensiveUcase creates a new instance of FiboRecursiveCPUIntensiveUcase
func NewFiboRecursiveCPUIntensiveUcase() contracts.CPUntensiveUcaseContract {
	return &FiboRecursiveCPUIntensiveUcase{}
}

// CPU-Intensive Function: Recursive Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// DoTask executes a CPU-intensive task synchronously
func (uc *FiboRecursiveCPUIntensiveUcase) DoTask(ctx context.Context, n int) int {
	return fibonacci(n)
}
