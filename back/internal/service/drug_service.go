package service

import (
	"context"
	"medical_farm/back/internal/models"
	"medical_farm/back/internal/repository"
)

type DrugService interface {
	ListAvailableDrugs(ctx context.Context, shopID int) ([]models.ShopInventory, error)
	GetDrugDetail(ctx context.Context, drugID int) (*models.Drug, error)
	// ... другие методы
}

type drugService struct {
	drugRepo repository.DrugRepository
}

func NewDrugService(drugRepo repository.DrugRepository) DrugService {
	return &drugService{drugRepo: drugRepo}
}

func (s *drugService) ListAvailableDrugs(ctx context.Context, shopID int) ([]models.ShopInventory, error) {
	// Здесь можно добавить бизнес-логику: фильтрация, сортировка, кэширование и т.д.
	return s.drugRepo.GetAvailableInShop(ctx, shopID)
}

func (s *drugService) GetDrugDetail(ctx context.Context, drugID int) (*models.Drug, error) {
	return s.drugRepo.GetByID(ctx, drugID)
}
