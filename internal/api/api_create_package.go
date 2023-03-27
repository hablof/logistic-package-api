package api

import (
	"context"

	"github.com/hablof/logistic-package-api/internal/model"
	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *logisticPackageAPI) CreatePackageV1(ctx context.Context, req *pb.CreatePackageV1Request) (*pb.CreatePackageV1Response, error) {

	log := o.setupLogger(ctx)

	log.Debug().Msg("logisticPackageAPI.CreatePackageV1 called")

	unit := model.Package{
		ID:            0, //
		Title:         req.GetTitle(),
		Material:      req.GetMaterial(),
		MaximumVolume: req.GetMaximumVolume(),
		Reusable:      req.GetReusable(),
	}

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("logisticPackageAPI.CreatePackageV1 failed")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	newID, err := o.repo.CreatePackage(ctx, &unit, log)
	if err != nil {
		log.Error().Err(err).Msg("repo.CreatePackage - failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	totalCUDevents.Inc()

	log.Debug().Msg("CreatePackageV1 - success")

	resp := pb.CreatePackageV1Response{
		ID: newID,
	}

	return &resp, nil
}