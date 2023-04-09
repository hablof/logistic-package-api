package service

import (
	"context"
	"fmt"

	"github.com/hablof/logistic-package-api/internal/model"

	"github.com/rs/zerolog"
)

type FieldName uint8

const (
	_ FieldName = iota
	Title
	Material
	MaxVolume
	Reusable
)

type RepoCRUD interface {
	CreatePackage(ctx context.Context, pack *model.Package, logger zerolog.Logger) (uint64, error)
	DescribePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64, limit uint64, logger zerolog.Logger) ([]model.Package, error)
	RemovePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) error
	UpdatePackage(ctx context.Context, packageID uint64, changes map[FieldName]interface{}, log zerolog.Logger) error
}

type ServiceCRUD interface {
	CreatePackage(ctx context.Context, pack *model.Package, logger zerolog.Logger) (uint64, error)
	DescribePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) (*model.Package, error)
	ListPackages(ctx context.Context, offset uint64, limit uint64, logger zerolog.Logger) ([]model.Package, error)
	RemovePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) error
	UpdatePackage(ctx context.Context, packageID uint64, changes map[FieldName]interface{}, log zerolog.Logger) error
}

var (
	ErrRepoEntityNotFound = fmt.Errorf("entity not found in repository")
)

type Service struct {
	r RepoCRUD
}

// CreatePackage implements api.ServiceCRUD
func (s *Service) CreatePackage(ctx context.Context, pack *model.Package, logger zerolog.Logger) (uint64, error) {
	return s.r.CreatePackage(ctx, pack, logger)
}

// DescribePackage implements api.ServiceCRUD
func (s *Service) DescribePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) (*model.Package, error) {
	return s.r.DescribePackage(ctx, packageID, logger)
}

// ListPackages implements api.ServiceCRUD
func (s *Service) ListPackages(ctx context.Context, offset uint64, limit uint64, logger zerolog.Logger) ([]model.Package, error) {
	return s.r.ListPackages(ctx, offset, limit, logger)
}

// RemovePackage implements api.ServiceCRUD
func (s *Service) RemovePackage(ctx context.Context, packageID uint64, logger zerolog.Logger) error {
	return s.r.RemovePackage(ctx, packageID, logger)
}

// UpdatePackage implements api.ServiceCRUD
func (s *Service) UpdatePackage(ctx context.Context, packageID uint64, changes map[FieldName]interface{}, log zerolog.Logger) error {
	return s.r.UpdatePackage(ctx, packageID, changes, log)
}

func NewService(r RepoCRUD) ServiceCRUD {
	return &Service{
		r: r,
	}
}
