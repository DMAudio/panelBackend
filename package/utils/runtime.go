package utils

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/DMAudio/panelBackend/package/log"
)

func RuntimeHandleQuitSignal(onQuit func()) {
	var quit = make(chan os.Signal)
	signal.Notify(
		quit,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	go func() {
		sig := <-quit
		log.Warnf("received signal `%s`, going to shutdown!", sig.String())
		if onQuit != nil {
			onQuit()
		}
		os.Exit(0)
	}()
}
