package sender

import (
	"github.com/hablof/logistic-package-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.PackageEvent) error
}
