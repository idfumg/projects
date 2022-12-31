package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"myapp/handlers"
	"myapp/repo"
	"myapp/services"
)

// func main() {
// 	app := fiber.New()
// 	app.Get("/hello", func(c *fiber.Ctx) error {
// 		// time.Sleep(10 * time.Millisecond)
// 		return c.SendString("hello world")
// 	})
// 	app.Listen(":8080")
// }

func main() {
	db, cancel := initDB()
	defer cancel()

	// redisClient := initRedis()

	productRepo := repo.NewProductRepoDB(db)
	// productRepo := repo.NewProductRepoRedis(db, redisClient)
	// productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	productService := services.NewCatalogService(productRepo)
	productHandler := handlers.NewCatalogHandler(productService)
	app := fiber.New()
	app.Get("/products", productHandler.GetProducts)
	app.Listen(":8080")
}

func initDB() (*gorm.DB, func()) {
	dsn := "root:password@tcp(localhost:3306)/infinitas"
	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	cancel := func() {
		sqldb, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqldb.Close()
	}
	return db, cancel
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
