package contracts

import "context"

type CPUntensiveUcaseContract interface {
	DoTask(ctx context.Context, n int) int
}
