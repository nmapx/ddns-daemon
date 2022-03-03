package cmd

import (
	elog "github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"syscall"
)

func HandleSignals() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigchan
		elog.Infof("Program killed")
		os.Exit(130)
	}()
}
