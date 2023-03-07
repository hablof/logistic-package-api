package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hablof/logistic-package-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := retranslator.Config{
		ChannelSize:     512,
		ConsumerCount:   2,
		BatchSize:       10,
		ConsumeInterval: 0,
		ProducerCount:   28,
		WorkerCount:     2,
		Repo:            nil,
		Sender:          nil,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	<-sigs

	retranslator.Close()
}
