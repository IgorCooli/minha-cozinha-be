package shoppingList

import (
	"context"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"github.com/IgorCooli/minha-cozinha-be/internal/repository/shoppingList"
	"github.com/google/uuid"
)

type Service interface {
	AddShoppingList(ctx context.Context, expense model.Item) error
	Search(ctx context.Context, name string) []model.Item
	RemoveShoppingListItem(ctx context.Context, id string) error
}

type service struct {
	repository shoppingList.Repository
}

func NewService(repository shoppingList.Repository) Service {
	return service{
		repository: repository,
	}
}

func (s service) AddShoppingList(ctx context.Context, shoppingListItem model.Item) error {

	shoppingListItem.Id = uuid.New().String()

	s.repository.InsertOne(ctx, shoppingListItem)

	return nil
}

func (s service) Search(ctx context.Context, name string) []model.Item {
	result := s.repository.Search(ctx, name)

	if result == nil {
		return []model.Item{}
	}

	return result
}

func (s service) RemoveShoppingListItem(ctx context.Context, id string) error {
	return s.repository.RemoveItem(ctx, id)
}
