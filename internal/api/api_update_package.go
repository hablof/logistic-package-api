package api

import (
	"context"
	"errors"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FieldName uint8

const (
	_ FieldName = iota
	Title
	Material
	MaxVolume
	Reusable
)

func (o *logisticPackageAPI) UpdatePackageV1(ctx context.Context, req *pb.UpdatePackageV1Request) (*pb.UpdatePackageV1Response, error) {

	log := o.setupLogger(ctx)

	log.Debug().Msg("logisticPackageAPI.UpdatePackageV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("logisticPackageAPI.UpdatePackageV1 failed")

		if err, ok := err.(pb.UpdatePackageV1RequestValidationError); ok {
			return nil, status.Error(codes.InvalidArgument, err.Field())
		}

		return nil, status.Error(codes.InvalidArgument, "unable to fetch invalid field")
	}

	changes := make(map[FieldName]interface{}, 4) // pb package struct has 4 non id fields

	if title := req.GetTitle(); title != "" {
		changes[Title] = title
	}

	if material := req.GetMaterial(); material != "" {
		changes[Material] = material
	}

	if maxVolume := req.GetMaximumVolume(); maxVolume != 0 {
		changes[MaxVolume] = maxVolume
	}

	if msgReuseable := req.GetReusable(); msgReuseable != nil {
		changes[Reusable] = msgReuseable.GetReusable()
	}

	if len(changes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no changes specified")
	}

	switch err := o.repo.UpdatePackage(ctx, req.GetPackageID(), changes, log); {
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

	resp := pb.UpdatePackageV1Response{
		Suc: true,
	}

	return &resp, nil
}
