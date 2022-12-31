package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/services"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type catalogHandlerRedis struct {
	catalogSrv  services.CatalogService
	redisClient *redis.Client
}

func NewCatalogHandlerRedis(catalogSrv services.CatalogService, redisClient *redis.Client) CatalogHandler {
	return catalogHandlerRedis{
		catalogSrv:  catalogSrv,
		redisClient: redisClient,
	}
}

func (h catalogHandlerRedis) GetProducts(c *fiber.Ctx) error {
	key := "handler::GetProducts"
	if responseJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("redis")
		c.Set("Content-Type", "application/json")
		return c.SendString(responseJson)
	}
	products, err := h.catalogSrv.GetProducts()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status":   "ok",
		"products": products,
	}
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), 0)
	}
	fmt.Println("database")
	return c.JSON(response)
}
