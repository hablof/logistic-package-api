package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/hablof/logistic-package-api/internal/repo"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "logistic_package_api_template_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type logisticPackageAPI struct {
	pb.UnimplementedLogisticPackageApiServiceServer
	repo repo.Repo
}

// NewLogisticPackageAPI returns api of logistic-package-api service
func NewLogisticPackageAPI(r repo.Repo) pb.LogisticPackageApiServiceServer {
	return &logisticPackageAPI{repo: r}
}
