package api

import (
	"context"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"
	"github.com/rs/zerolog"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *logisticPackageAPI) ListPackagesV1(ctx context.Context, req *pb.ListPackagesV1Request) (*pb.ListPackagesV1Response, error) {

	log := o.logger
	if o.shouldRiseDebugLevel(ctx) {
		log = log.Level(zerolog.DebugLevel)
	}

	log.Debug().Msg("logisticPackageAPI.ListPackagesV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("logisticPackageAPI.ListPackageV1 failed")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	packageList, err := o.repo.ListPackages(ctx, req.GetOffset(), log)
	if err != nil {
		log.Error().Err(err).Msg("repo.ListPackage - failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListPackageV1 - success")

	output := make([]string, len(packageList))
	for i, pack := range packageList {
		output[i] = pack.Title
	}

	resp := pb.ListPackagesV1Response{
		PackageTitle: output,
	}

	return &resp, nil
}
