package main

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OTP struct { // one time password
	Key     string
	Created time.Time
}

type RetensionMap map[string]OTP

func NewRetentionMap(ctx context.Context, retensionPeriod time.Duration) RetensionMap {
	ans := RetensionMap{}
	go ans.Retention(ctx, retensionPeriod)
	return ans
}

func (r RetensionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}
	r[o.Key] = o
	return o
}

func (r RetensionMap) VerifyOTP(otp string) bool {
	if _, ok := r[otp]; !ok {
		return false
	}
	delete(r, otp)
	return true
}

func (r RetensionMap) Retention(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, otp := range r {
				if otp.Created.Add(retentionPeriod).Before(time.Now()) {
					delete(r, otp.Key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
