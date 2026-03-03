package service

import (
	"context"

	"github.com/summit/summit-api/internal/models"
	"github.com/summit/summit-api/internal/repository"
)

type RegionService struct {
	regionRepo *repository.RegionRepository
}

func NewRegionService(rr *repository.RegionRepository) *RegionService {
	return &RegionService{regionRepo: rr}
}

func (s *RegionService) List(ctx context.Context) ([]models.Region, error) {
	return s.regionRepo.List(ctx)
}
