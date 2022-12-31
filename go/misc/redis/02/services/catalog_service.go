package services

import "myapp/repo"

type catalogService struct {
	productRepo repo.ProductRepo
}

func NewCatalogService(productRepo repo.ProductRepo) CatalogService {
	return catalogService{
		productRepo: productRepo,
	}
}

func (s catalogService) GetProducts() (products []product, err error) {
	productsDB, err := s.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	for _, p := range productsDB {
		products = append(products, product{
			ID: p.ID,
			Name: p.Name,
			Quantity: p.Quantity,
		})
	}
	return products, nil
}