package mocks

import (
	"time"

	model "github.com/hablof/logistic-package-api/internal/model"
	"github.com/rs/zerolog/log"
)

type SillySender struct {
}

// Send mocks base method.
func (m *SillySender) Send(unit *model.PackageEvent) error {
	log.Debug().Msgf("silly sender called on event id %d", unit.ID)
	time.Sleep(10 * time.Second)
	// if unit.ID%2 == 0 {
	// 	return errors.New("sending failed")
	// }

	return nil
}
