package repo

type product struct {
	ID       int
	Name     string
	Quantity int
}

type ProductRepo interface {
	GetProducts() ([]product, error)
}


