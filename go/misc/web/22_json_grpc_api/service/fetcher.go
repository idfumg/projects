package service

import (
	"context"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}
