package stock

import (
	"context"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"github.com/IgorCooli/minha-cozinha-be/internal/repository/stock"
)

type Service interface {
	AddStock(ctx context.Context, expense model.StockItem) error
	Search(ctx context.Context, name string) []model.StockItem
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

	s.repository.InsertOne(ctx, stockItem)

	return nil
}

func (s service) Search(ctx context.Context, name string) []model.StockItem {

	return s.repository.Search(ctx, name)
}
