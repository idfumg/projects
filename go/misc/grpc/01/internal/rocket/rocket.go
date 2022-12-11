//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket myapp/internal/rocket Store

package rocket

import "context"

type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

type Store interface {
	GetRocket(id string) (Rocket, error)
	AddRocket(rocket Rocket) (Rocket, error)
	DelRocket(id string) error
}

type Service struct {
	Store Store
}

func NewService(store Store) (Service, error) {
	return Service{
		Store: store,
	}, nil
}

func (s Service) GetRocket(ctx context.Context, id string) (Rocket, error) {
	rocket, err := s.Store.GetRocket(id)
	if err != nil {
		return Rocket{}, err
	}
	return rocket, nil
}

func (s Service) AddRocket(ctx context.Context, rocket Rocket) (Rocket, error) {
	rocket, err := s.Store.AddRocket(rocket)
	if err != nil {
		return Rocket{}, err
	}
	return rocket, nil
}

func (s Service) DelRocket(ctx context.Context, id string) error {
	err := s.Store.DelRocket(id)
	if err != nil {
		return err
	}
	return nil
}
