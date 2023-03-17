package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/hablof/logistic-package-api/internal/model"
	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "logistic_package_api_template_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type RepoCRUD interface {
	CreatePackage(ctx context.Context, pack *model.Package) (uint64, error)
	DescribePackage(ctx context.Context, packageID uint64) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64) ([]*model.Package, error)
	RemovePackage(ctx context.Context, packageID uint64) error
}

type logisticPackageAPI struct {
	pb.UnimplementedLogisticPackageApiServiceServer
	repo RepoCRUD
}

// NewLogisticPackageAPI returns api of logistic-package-api service
func NewLogisticPackageAPI(r RepoCRUD) pb.LogisticPackageApiServiceServer {
	return &logisticPackageAPI{repo: r}
}
