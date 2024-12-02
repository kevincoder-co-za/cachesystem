package core

import (
	"time"
)

func cleanupExpiredEntries() {
	CacheStore.Range(func(key, value interface{}) bool {
		if entry, ok := value.(CacheEntry); ok {
			if entry.ExpiresOn.Before(time.Now()) {
				CacheStore.Delete(key)
			}
		}
		return true
	})
}

func RunCrons() {
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			cleanupExpiredEntries()
		}
	}
}
