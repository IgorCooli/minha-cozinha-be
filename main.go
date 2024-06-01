package main

import (
	"context"
	"os"
	"time"

	shoppingListHandler "github.com/IgorCooli/minha-cozinha-be/api/shoppingList"
	stockHandler "github.com/IgorCooli/minha-cozinha-be/api/stock"
	shoppingListService "github.com/IgorCooli/minha-cozinha-be/internal/business/service/shoppingList"
	stockService "github.com/IgorCooli/minha-cozinha-be/internal/business/service/stock"
	shoppingListRepository "github.com/IgorCooli/minha-cozinha-be/internal/repository/shoppingList"
	stockRepository "github.com/IgorCooli/minha-cozinha-be/internal/repository/stock"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dbClient := setupDb(ctx)

	app := fiber.New()
	app.Use(cors.New(cors.Config{

		AllowOrigins:     "*",
		AllowCredentials: false,
	}))

	injectStockApi(ctx, dbClient, app)
	injectShoppingListApi(ctx, dbClient, app)

	port := resolveApiPort()

	app.Listen(":" + port)
}

func setupDb(ctx context.Context) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:mongodb159@tccmongodb.3ud5x.mongodb.net/?retryWrites=true&w=majority&appName=TCCMongoDB"))
	if err != nil {
		panic("Could not connect to dabase")
	}

	return client
}

func injectStockApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {

	stockRepository := stockRepository.NewRepository(dbClient)
	stockService := stockService.NewService(stockRepository)
	stockHandler.NewHandler(ctx, stockService, app)
}

func injectShoppingListApi(ctx context.Context, dbClient *mongo.Client, app *fiber.App) {

	shoppingListRepository := shoppingListRepository.NewRepository(dbClient)
	shoppingListService := shoppingListService.NewService(shoppingListRepository)
	shoppingListHandler.NewHandler(ctx, shoppingListService, app)
}

func resolveApiPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}
