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
	mutex  sync.Mutex
	once   sync.Once
)

func Ctx() context.Context {

	mutex.Lock()
	defer mutex.Unlock()
	once.Do(func() {
		ctx, cancel = context.WithCancel(context.Background())
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

func Force() {
	mutex.Lock()
	defer mutex.Unlock()
	if cancel != nil {
		cancel()
	}
}
