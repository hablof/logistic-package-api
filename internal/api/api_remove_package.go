package api

import (
	"context"
	"errors"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *logisticPackageAPI) RemovePackageV1(ctx context.Context, req *pb.RemovePackageV1Request) (*pb.RemovePackageV1Response, error) {

	log := o.setupLogger(ctx)

	log.Debug().Msg("logisticPackageAPI.RemovePackageV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("RemovePackageV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	switch err := o.repo.RemovePackage(ctx, req.GetPackageID(), log); {
	case errors.Is(err, ErrRepoEntityNotFound):
		log.Debug().Uint64("packageID", req.PackageID).Msg("package not found")
		totalTemplateNotFound.Inc()

		return nil, status.Error(codes.NotFound, "package not found")

	case err != nil:
		log.Error().Err(err).Msg("RemovePackageV1 - failed")
		return nil, status.Error(codes.Internal, err.Error())

	}

	totalCUDevents.Inc()

	log.Debug().Msg("RemovePackageV1 - success")

	resp := pb.RemovePackageV1Response{
		Suc: true,
	}

	return &resp, nil
}
