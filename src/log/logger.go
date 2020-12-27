package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Logger is the global logger.
var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return Logger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return Logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return Logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return Logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return Logger.Hook(h)
}

// Err starts a new message with error level with err as a field if not nil or
// with info level if err is nil.
//
// You must call Msg on the returned event in order to send the event.
func Err(err error) *zerolog.Event {
	return Logger.Err(err)
}

func ErrC(ctx context.Context, err error) *zerolog.Event {
	return Logger.Err(err)
}

func ErrLT(err error) *zerolog.Event {
	return Logger.Err(err).Str(longTermKey, longTermTrue)
}

func ErrCLT(ctx context.Context, err error) *zerolog.Event {
	return Ctx(ctx).Err(err).Str(longTermKey, longTermTrue)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func Trace() *zerolog.Event {
	return Logger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return Logger.Debug()
}

func DebugC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Debug()
}

func DebugLT(err error) *zerolog.Event {
	return Logger.Debug().Str(longTermKey, longTermTrue)
}

func DebugCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Debug().Str(longTermKey, longTermTrue)
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return Logger.Info()
}

func InfoC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Info()
}

func InfoLT() *zerolog.Event {
	return Logger.Info().Str(longTermKey, longTermTrue)
}

func InfoCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Info().Str(longTermKey, longTermTrue)
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return Logger.Warn()
}

func WarnC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Warn()
}

func WarnLT() *zerolog.Event {
	return Logger.Warn().Str(longTermKey, longTermTrue)
}

func WarnCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Warn().Str(longTermKey, longTermTrue)
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return Logger.Error()
}

func ErrorC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Error()
}

func ErrorLT() *zerolog.Event {
	return Logger.Error().Str(longTermKey, longTermTrue)
}

func ErrorCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Error().Str(longTermKey, longTermTrue)
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

func FatalC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Fatal()
}

func FatalLT() *zerolog.Event {
	return Logger.Fatal().Str(longTermKey, longTermTrue)
}

func FatalCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Fatal().Str(longTermKey, longTermTrue)
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return Logger.Panic()
}

func PanicC(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Panic()
}

func PanicLT() *zerolog.Event {
	return Logger.Panic().Str(longTermKey, longTermTrue)
}

func PanicCLT(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Panic().Str(longTermKey, longTermTrue)
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return Logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return Logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	logger, ok := ctx.Value(ctxLoggerKey).(*zerolog.Logger)
	if !ok || logger == nil {
		return &Logger
	}
	return logger
}
