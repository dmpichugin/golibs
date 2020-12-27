package log

const (
	longTermKey  = "lt"
	longTermTrue = "t"
)

type contextKey uint64

const (
	ctxLoggerKey contextKey = iota
)
