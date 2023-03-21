package limiter

import (
	"errors"
	"sync"
)

const (
	Fixed   = 1
	Sliding = 2
)

type RateLimiter struct {
	limiterType int8
	limiter     InterfaceLimiter
	mu          *sync.RWMutex
	duration    uint32
	count       uint32
}

// NewRateLimiter 创建一个RateLimiter
func NewRateLimiter(limiterType int8, duration, count uint32) (*RateLimiter, error) {
	var (
		newLimiter InterfaceLimiter
		err        error
	)

	switch limiterType {
	case Fixed:
		newLimiter, err = NewFixedLimiter(duration, count)
	// case Sliding:
	// 	newLimiter, err = limiter.NewSlidingLimiter(duration, count)
	default:
		return nil, errors.New("limiter type is invalid")
	}
	if err != nil {
		return nil, err
	}
	return &RateLimiter{
		limiterType: limiterType,
		limiter:     newLimiter,
		mu:          &sync.RWMutex{},
		duration:    duration,
		count:       count,
	}, nil
}

// Allow 检查是否通过
func (l *RateLimiter) Allow(id int32) bool {
	return l.limiter.Check(id)
}

// AllowNum 根据Num检查是否通过
func (l *RateLimiter) AllowNum(id int32, num uint32) bool {
	return l.limiter.CheckNum(id, num)
}
