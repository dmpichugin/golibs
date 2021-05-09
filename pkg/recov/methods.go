package recov

import (
	"context"
)

var (
	R Recoverer = NewZerologRecoverer()
)

func Err(ctx context.Context) {
	r := recover()
	if r != nil {
		R.Err(ctx, r)
	}
}
