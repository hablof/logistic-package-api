package api

import (
	"context"

	pb "github.com/hablof/logistic-package-api/pkg/logistic-package-api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (o *logisticPackageAPI) ListPackagesV1(ctx context.Context, req *pb.ListPackagesV1Request) (*pb.ListPackagesV1Response, error) {

	log := o.setupLogger(ctx)

	log.Debug().Msg("logisticPackageAPI.ListPackagesV1 called")

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("logisticPackageAPI.ListPackageV1 failed")

		if err, ok := err.(pb.ListPackagesV1RequestValidationError); ok {
			return nil, status.Error(codes.InvalidArgument, err.Field())
		}

		return nil, status.Error(codes.InvalidArgument, "unable to fetch invalid field")
	}

	packageList, err := o.service.ListPackages(ctx, req.GetOffset(), req.GetLimit(), log)
	if err != nil {
		log.Error().Err(err).Msg("repo.ListPackage - failed")
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Debug().Msg("ListPackageV1 - success")

	output := make([]*pb.Package, len(packageList))
	for i, pack := range packageList {
		unit := &pb.Package{
			ID:            pack.ID,
			Title:         pack.Title,
			Material:      pack.Material,
			MaximumVolume: pack.MaximumVolume,
			Reusable:      pack.Reusable,
			Created:       timestamppb.New(pack.Created),
			Updated:       nil,
		}
		if pack.Updated != nil {
			unit.Updated = &pb.MaybeTimestamp{
				Time: timestamppb.New(*pack.Updated),
			}
		}

		output[i] = unit
	}

	resp := pb.ListPackagesV1Response{
		PackageTitle: nil,
		Packages:     output,
	}

	return &resp, nil
}
