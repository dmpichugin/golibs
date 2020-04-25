package shutdown

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	once sync.Once
)

func Ctx() context.Context {

	once.Do(func() {
		initCtx()
		go listener()
	})

	return ctx
}

func listener() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
	case <-signalChan:
	}
	signal.Stop(signalChan)

	cancel()
}

func initCtx() {
	ctx, cancel = context.WithCancel(context.Background())
}

func Force() {
	if cancel != nil {
		cancel()
	}
}
