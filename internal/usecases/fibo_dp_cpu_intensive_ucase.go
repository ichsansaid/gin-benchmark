package ucases

import (
	"context"

	"github.com/ichsansaid/golang-gin-benchmark/internal/contracts"
)

// FiboDPCPUIntensiveUcase implements FiboDPCPUIntensiveUcaseContract
type FiboDPCPUIntensiveUcase struct{}

// NewFiboDPCPUIntensiveUcase creates a new instance of FiboDPCPUIntensiveUcase
func NewFiboDPCPUIntensiveUcase() contracts.CPUntensiveUcaseContract {
	return &FiboDPCPUIntensiveUcase{}
}

// CPU-Intensive Function: Iterative Fibonacci using dynamic programming
func fibonacciStack(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// DoTask executes a CPU-intensive task synchronously
func (uc *FiboDPCPUIntensiveUcase) DoTask(ctx context.Context, n int) int {
	return fibonacciStack(n)
}
