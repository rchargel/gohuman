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
	items      map[string]timedItem
	lastPurge  time.Time
	lock       *sync.Mutex
	wg         sync.WaitGroup
	deleteChan chan string
}

var store *timedStore

func init() {
	store = newStore()
}

func newStore() *timedStore {
	s := &timedStore{
		items:      make(map[string]timedItem, 100),
		lastPurge:  time.Now(),
		lock:       &sync.Mutex{},
		wg:         sync.WaitGroup{},
		deleteChan: make(chan string),
	}
	go s.deleteKeysAsNeeded()
	return s
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
	s.lock.Unlock()
	if sinceLastCheck > maxPurgeCheckSeconds {
		// start the WaitGroup
		s.wg.Add(1)
		go s.findKeysToRemove(maxSeconds)
		s.lock.Lock()
		s.lastPurge = time.Now()
		s.lock.Unlock()
	}
}

func (s *timedStore) findKeysToRemove(maxSeconds int) {
	for key, value := range s.items {
		since := int(time.Since(value.when).Seconds())
		if since >= maxSeconds {
			s.wg.Add(1)
			s.deleteChan <- key
		}
	}
	s.wg.Done()
}

func (s *timedStore) waitForPurgeToComplete() {
	s.wg.Wait()
}

func (s *timedStore) deleteKeysAsNeeded() {
	for key := range s.deleteChan {
		s.lock.Lock()
		//fmt.Println("Delete: " + key)
		delete(s.items, key)
		s.wg.Done()
		s.lock.Unlock()
	}
}
