package stock

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/IgorCooli/minha-cozinha-be/internal/business/model"
	"github.com/IgorCooli/minha-cozinha-be/internal/business/service/stock"
	"github.com/gofiber/fiber/v3"
)

type handler struct {
	service stock.Service
}

func NewHandler(ctx context.Context, service stock.Service, app *fiber.App) handler {

	handler := handler{
		service: service,
	}

	app.Get("/", handler.HelloWorld)
	app.Get("/stock/search", handler.SearchStock)
	app.Post("/stock", handler.AddStock)
	app.Delete("/stock/:id", handler.RemoveStockItem)

	return handler
}

func (h handler) HelloWorld(c fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", c.Params("*"))
	err := c.SendString(msg) // => ✋ register

	if err != nil {
		panic("")
	}

	return nil
}

func (h handler) SearchStock(c fiber.Ctx) error {

	name := c.Query("name")

	result := h.service.Search(c.Context(), name)

	c.JSON(result)
	return nil
}

func (h handler) AddStock(c fiber.Ctx) error {
	var body model.StockItem
	json.Unmarshal(c.Body(), &body)

	h.service.AddStock(c.Context(), body)

	return nil
}

func (h handler) RemoveStockItem(c fiber.Ctx) error {

	stockItemId := c.Params("id")

	h.service.RemoveStockItem(c.Context(), stockItemId)

	return nil
}
