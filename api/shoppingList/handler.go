package shoppingList

import (
	"context"
	"encoding/json"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"github.com/IgorCooli/minha-cozinha-be/internal/business/service/shoppingList"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service shoppingList.Service
}

func NewHandler(ctx context.Context, service shoppingList.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Get("/shopping-list/search", handler.SearchShoppingList)
	app.Post("/shopping-list", handler.AddShoppingList)
	app.Delete("/shopping-list/:id", handler.RemoveShoppingListItem)

	return handler
}

func (h handler) SearchShoppingList(c fiber.Ctx) error {

	name := c.Query("name")

	result := h.service.Search(c.Context(), name)

	c.JSON(result)
	return nil
}

func (h handler) AddShoppingList(c fiber.Ctx) error {
	var body model.Item
	json.Unmarshal(c.Body(), &body)

	h.service.AddShoppingList(c.Context(), body)

	return nil
}

func (h handler) RemoveShoppingListItem(c fiber.Ctx) error {

	shoppingListItemId := c.Params("id")

	h.service.RemoveShoppingListItem(c.Context(), shoppingListItemId)

	return nil
}
