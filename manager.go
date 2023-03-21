package limiter

import (
	"errors"
	"sync"
)

var BaseManagerLimiter *ManagerLimiter

func init() {
	BaseManagerLimiter = NewLimiterManager()
}

func BaseAddLimiter(name string, limiterType int8, duration, count uint32) error {
	return BaseManagerLimiter.AddLimiter(name, limiterType, duration, count)
}

func BaseAllow(name string, id int32) (bool, error) {
	return BaseManagerLimiter.Allow(name, id)
}

func BaseAllowNum(name string, id int32, num uint32) (bool, error) {
	return BaseManagerLimiter.AllowNum(name, id, num)
}

type ManagerLimiter struct {
	m  map[string]*RateLimiter
	mu *sync.RWMutex
}

func NewLimiterManager() *ManagerLimiter {
	return &ManagerLimiter{
		m:  make(map[string]*RateLimiter),
		mu: &sync.RWMutex{},
	}
}

func (m *ManagerLimiter) AddLimiter(name string, limiterType int8, duration, count uint32) error {
	if len(name) == 0 {
		return errors.New("name is nil")
	}
	if _, ok := m.m[name]; ok {
		return errors.New("name is already existed")
	}
	if duration > 86400*30 {
		return errors.New("duration max 86400*30")
	}
	if count > 1000000000 {
		return errors.New("count max 1000000000")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	l, err := NewRateLimiter(limiterType, duration, count)
	if err != nil {
		return err
	}
	m.m[name] = l
	return nil
}

func (m *ManagerLimiter) Allow(name string, id int32) (bool, error) {
	if _, ok := m.m[name]; !ok {
		return false, errors.New("name is not existed")
	}
	return m.m[name].Allow(id), nil
}

func (m *ManagerLimiter) AllowNum(name string, id int32, num uint32) (bool, error) {
	if _, ok := m.m[name]; !ok {
		return false, errors.New("name is not existed")
	}
	return m.m[name].AllowNum(id, num), nil
}
