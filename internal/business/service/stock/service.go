package stock

import (
	"context"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"github.com/IgorCooli/minha-cozinha-be/internal/repository/stock"
	"github.com/google/uuid"
)

type Service interface {
	AddStock(ctx context.Context, expense model.StockItem) error
	Search(ctx context.Context, name string) []model.StockItem
	RemoveStockItem(ctx context.Context, id string) error
}

type service struct {
	repository stock.Repository
}

func NewService(repository stock.Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) AddStock(ctx context.Context, stockItem model.StockItem) error {

	stockItem.Id = uuid.New().String()

	s.repository.InsertOne(ctx, stockItem)

	return nil
}

func (s service) Search(ctx context.Context, name string) []model.StockItem {
	result := s.repository.Search(ctx, name)

	if result == nil {
		return []model.StockItem{}
	}

	return result
}

func (s service) RemoveStockItem(ctx context.Context, id string) error {
	return s.repository.RemoveItem(ctx, id)
}
