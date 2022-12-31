package service

import (
	"context"
	"fmt"
	"myapp/config"
	"time"
)

var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG":  100_000.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond) // mimic the http roundtrip

	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}
	return price, nil
}

type priceFetcher struct {
	cfg *config.Cfg
}

func NewPriceFetcher(cfg *config.Cfg) *priceFetcher {
	return &priceFetcher{
		cfg: cfg,
	}
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}
