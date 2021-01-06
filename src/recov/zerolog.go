package recov

import (
	"context"

	"github.com/dmpichugin/golibs/src/errors"
	"github.com/dmpichugin/golibs/src/log"
)

type zerologRecoverer struct{}

func NewZerologRecoverer() *zerologRecoverer {
	return &zerologRecoverer{}
}

func (r *zerologRecoverer) Err(ctx context.Context, recoverMessage interface{}) {
	if recoverMessage != nil {
		log.ErrorC(ctx).Stack().Err(errors.Errorf("%s", recoverMessage)).Send()
	}
}
