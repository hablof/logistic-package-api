package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hablof/logistic-package-api/internal/app/retranslator"
	"github.com/hablof/logistic-package-api/internal/mocks"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := retranslator.RetranslatorConfig{
		ChannelSize:     512,
		ConsumerCount:   2,
		BatchSize:       10,
		ConsumeInterval: 2 * time.Second,
		ProducerCount:   28,
		WorkerCount:     2,
		Repo:            &mocks.MockEventRepo{},
		Sender:          &mocks.MockEventSender{},
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	<-sigs

	retranslator.Close()
}
