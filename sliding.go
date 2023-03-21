package limiter

import (
	"errors"
	"sync"
	"time"
)

type SlidingLimiter struct {
	m        map[int32]uint32
	mu       *sync.RWMutex
	duration uint32
	count    uint32
}

func NewSlidingLimiter(duration, count uint32) (*SlidingLimiter, error) {
	if duration == 0 {
		return nil, errors.New("duration has to be greater than 0")
	}
	if count == 0 {
		return nil, errors.New("count has to be greater than 0")
	}
	s := &SlidingLimiter{
		m:        make(map[int32]uint32),
		mu:       &sync.RWMutex{},
		duration: duration,
		count:    count,
	}
	go func() {
		limit := time.NewTicker(time.Second * time.Duration(duration))
		select {
		case <-limit.C:
			s.reset()
		}
	}()
	return s, nil
}

func (l *SlidingLimiter) reset() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.m = make(map[int32]uint32)
}

func (l *SlidingLimiter) Check(id int32) bool {
	l.mu.Lock()
	if _, ok := l.m[id]; !ok {
		l.m[id] = 0
	}
	l.m[id]++
	n := l.m[id]
	l.mu.Unlock()
	if n > l.count {
		return false
	}
	return true
}
