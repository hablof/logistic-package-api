package api

import (
	"context"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *logisticPackageAPI) RemovePackageV1(ctx context.Context, req *pb.RemovePackageV1Request) (*pb.RemovePackageV1Response, error) {
	log.Debug().Msg("logisticPackageAPI.RemovePackageV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemovePackageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := o.repo.RemovePackage(ctx, req.GetPackageID())
	if err != nil {
		log.Error().Err(err).Msg("RemovePackageV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("RemovePackageV1 - success")

	resp := pb.RemovePackageV1Response{
		Suc: err == nil,
	}

	return &resp, nil
}