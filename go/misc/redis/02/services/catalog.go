package services

type product struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type CatalogService interface {
	GetProducts() ([]product, error)
}
