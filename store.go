package gohuman

import (
	"fmt"
	"sync"
	"time"
)

const maxTimeToLiveMinutes = 20
const maxTimeBeforePurgeSeconds = 100

type timedItem struct {
	when time.Time
	item Captcha
}

type timedStore struct {
	items     map[string]timedItem
	lastPurge time.Time
	lock      *sync.Mutex
}

var store *timedStore

func init() {
	store = newStore()
}

func newStore() *timedStore {
	return &timedStore{
		items:     make(map[string]timedItem, 100),
		lastPurge: time.Now(),
		lock:      &sync.Mutex{},
	}
}

func (s *timedStore) addCaptcha(captcha Captcha) {
	s.purgeOld(maxTimeToLiveMinutes*60, maxTimeBeforePurgeSeconds)
	s.items[captcha.ID] = timedItem{
		when: time.Now(),
		item: captcha,
	}
}

func (s *timedStore) getCaptcha(id string) (Captcha, error) {
	if c, ok := s.items[id]; ok {
		return c.item, nil
	}
	return Captcha{}, fmt.Errorf("No captcha found with ID: %v", id)
}

func (s *timedStore) purgeOld(maxSeconds, maxPurgeCheckSeconds int) {
	s.lock.Lock()
	sinceLastCheck := int(time.Since(s.lastPurge).Seconds())
	if sinceLastCheck > maxPurgeCheckSeconds {
		for key, value := range s.items {
			since := int(time.Since(value.when).Seconds())
			if since >= maxSeconds {
				delete(s.items, key)
			}
		}
		s.lastPurge = time.Now()
	}
	s.lock.Unlock()
}
