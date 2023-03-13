package api

import (
	"context"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (o *logisticPackageAPI) DescribePackageV1(ctx context.Context, req *pb.DescribePackageV1Request) (*pb.DescribePackageV1Response, error) {

	log.Debug().Msg("logisticPackageAPI.DescribePackageV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribePackageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mypackage, err := o.repo.DescribePackage(ctx, req.GetPackageID())
	if err != nil {
		log.Error().Err(err).Msg("DescribePackageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if mypackage == nil {
		log.Debug().Uint64("packageID", req.PackageID).Msg("package not found")
		totalTemplateNotFound.Inc()

		return nil, status.Error(codes.NotFound, "package not found")
	}

	log.Debug().Msg("DescribePackageV1 - success")

	return &pb.DescribePackageV1Response{
		Value: &pb.Package{
			ID:            mypackage.ID,
			Title:         mypackage.Title,
			Material:      mypackage.Material,
			MaximumVolume: mypackage.MaximumVolume,
			Reusable:      mypackage.Reusable,
			Created: &timestamppb.Timestamp{
				Seconds: 0,
				Nanos:   0,
			},
		},
	}, nil
}
