package service

import (
	"context"
	"time"

	"myapp/config"
	"myapp/types"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
	cfg  *config.Cfg
}

func NewWithLogsService(next PriceFetcher, cfg *config.Cfg) PriceFetcher {
	return &loggingService{
		next: next,
		cfg:  cfg,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value(types.RequestIdKey("requestID")),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
