package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

const (
	longTermKey  = "lt"
	longTermTrue = "t"
)

type contextKey uint64

const (
	ctxLoggerKey contextKey = iota
)

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}
