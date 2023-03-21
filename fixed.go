package limiter

import (
	"errors"
	"sync"
	"time"
)

type FixedLimiter struct {
	m        map[int32]uint32
	mu       *sync.RWMutex
	duration uint32
	count    uint32
}

func NewFixedLimiter(duration, count uint32) (*FixedLimiter, error) {
	if duration == 0 {
		return nil, errors.New("duration has to be greater than 0")
	}
	if count == 0 {
		return nil, errors.New("count has to be greater than 0")
	}
	f := &FixedLimiter{
		m:        make(map[int32]uint32),
		mu:       &sync.RWMutex{},
		duration: duration,
		count:    count,
	}
	go func() {
		limit := time.NewTicker(time.Second * time.Duration(duration))
		select {
		case <-limit.C:
			f.reset()
		}
	}()
	return f, nil
}

func (l *FixedLimiter) reset() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.m = make(map[int32]uint32)
}

func (l *FixedLimiter) Check(id int32) bool {
	return l.CheckNum(id, 1)
}

func (l *FixedLimiter) CheckNum(id int32, num uint32) bool {
	if num > l.count {
		return false
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.m[id]; !ok {
		l.m[id] = num
		return true
	}
	n := l.m[id] + num
	if n > l.count {
		return false
	}
	l.m[id] = n
	return true
}
