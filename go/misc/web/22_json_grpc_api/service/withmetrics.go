package service

import (
	"context"
	"fmt"
	"myapp/config"
)

type metricService struct {
	next PriceFetcher
	cfg  *config.Cfg
}

func NewWithMetricsService(next PriceFetcher, cfg *config.Cfg) PriceFetcher {
	return &metricService{
		next: next,
		cfg:  cfg,
	}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("push to your metrics service")
	return s.next.FetchPrice(ctx, ticker)
}
