package api

import (
	"context"
	"errors"

	"github.com/hablof/logistic-package-api/internal/service"
	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	changes := make(map[service.FieldName]interface{}, 4) // pb package struct has 4 non id fields

	if title := req.GetTitle(); title != "" {
		changes[service.Title] = title
	}

	if material := req.GetMaterial(); material != "" {
		changes[service.Material] = material
	}

	if maxVolume := req.GetMaximumVolume(); maxVolume != 0 {
		changes[service.MaxVolume] = maxVolume
	}

	if msgReuseable := req.GetReusable(); msgReuseable != nil {
		changes[service.Reusable] = msgReuseable.GetReusable()
	}

	if len(changes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no changes specified")
	}

	switch err := o.service.UpdatePackage(ctx, req.GetPackageID(), changes, log); {
	case errors.Is(err, service.ErrRepoEntityNotFound):
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
