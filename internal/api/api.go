package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hablof/logistic-package-api/internal/repo"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "logistic_package_api_template_not_found_total",
		Help: "Total number of templates that were not found",
	})
)

type templateAPI struct {
	pb.UnimplementedLogisticPackageApiServiceServer
	repo repo.Repo
}

// NewTemplateAPI returns api of logistic-package-api service
func NewTemplateAPI(r repo.Repo) pb.LogisticPackageApiServiceServer {
	return &templateAPI{repo: r}
}

func (o *templateAPI) DescribeTemplateV1(
	ctx context.Context,
	req *pb.DescribePackageV1Request,
) (*pb.DescribePackageV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mypackage, err := o.repo.DescribeTemplate(ctx, req.PackageId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if mypackage == nil {
		log.Debug().Uint64("templateId", req.PackageId).Msg("template not found")
		totalTemplateNotFound.Inc()

		return nil, status.Error(codes.NotFound, "template not found")
	}

	log.Debug().Msg("DescribeTemplateV1 - success")

	return &pb.DescribePackageV1Response{
		Value: &pb.Package{
			Id:  mypackage.ID,
			Foo: mypackage.Foo,
		},
	}, nil
}
