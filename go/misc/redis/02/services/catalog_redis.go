package services

import (
	"context"
	"encoding/json"
	"fmt"
	"myapp/repo"

	"github.com/go-redis/redis/v8"
)

type catalogServiceRedis struct {
	productRepo repo.ProductRepo
	redisClient *redis.Client
}

func NewCatalogServiceRedis(productRepo repo.ProductRepo, redisClient *redis.Client) CatalogService {
	return catalogServiceRedis{
		productRepo: productRepo,
		redisClient: redisClient,
	}
}

func (s catalogServiceRedis) GetProducts() (products []product, err error) {
	key := "service::GetProducts"
	productJson, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		if json.Unmarshal([]byte(productJson), &products) == nil {
			fmt.Println("redis")
			return products, nil
		}
	}
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range productsDB {
		products = append(products, product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	if data, err := json.Marshal(products); err == nil {
		if s.redisClient.Set(context.Background(), key, data, 0).Err() == nil {
			return products, nil
		}
	}
	return products, nil
}
