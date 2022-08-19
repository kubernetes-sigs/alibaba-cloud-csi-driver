package auth

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"sync"
	"sync/atomic"
	"time"
)

const (
	refreshPeriod = 3 * time.Minute
)

var s *singleton

func init() {
	s = newSingleton()
}

type Empty struct{}

type Singleton interface {
	refresh()
}

type singleton struct {
	Once          sync.Once
	Value         atomic.Value
	StopC         chan struct{}
	RefreshPeriod time.Duration
	lock          sync.RWMutex
}

func newSingleton() *singleton {
	s := &singleton{
		Once:          sync.Once{},
		StopC:         make(chan struct{}),
		RefreshPeriod: refreshPeriod,
	}
	s.Once.Do(func() {
		s.refresh()
	})
	return s
}

func Get() *Empty {
	return s.Value.Load().(*Empty)
}

func (s *singleton) refresh() {
	go wait.Until(s.refreshVal, s.RefreshPeriod, s.StopC)
}

func (s *singleton) refreshVal() {
	o := &Empty{}
	s.Value.Store(o)
}
