package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepoRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepoRedis(db *gorm.DB, redisClient *redis.Client) ProductRepo {
	db.AutoMigrate(&product{})
	mockData(db)
	return productRepoRedis{
		db:          db,
		redisClient: redisClient,
	}
}

func (r productRepoRedis) GetProducts() (products []product, err error) {
	key := "repository::GetProducts"
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("redis")
			return products, nil
		}
	}
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), 10*time.Second).Err()
	if err != nil {
		return nil, err
	}
	fmt.Println("database")
	return products, nil
}
