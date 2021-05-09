package recov

import "context"

type Recoverer interface {
	Err(ctx context.Context, recoverMessage interface{})
}
