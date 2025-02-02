package contracts

import "context"

type IOIntensiveUcaseContract interface {
	DoTask(ctx context.Context)
}
