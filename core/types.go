package core

import (
	"sync"
	"time"
)

type ApiStatus string

var CacheStore sync.Map

const (
	StatusOK     ApiStatus = "success"
	StatusFailed ApiStatus = "failed"
)

type CacheEntry struct {
	Name      string      `json:"name"`
	ExpiresOn time.Time   `json:"expires"`
	Payload   interface{} `json:"payload"`
}

type StandardApiResponse struct {
	Status ApiStatus   `json:"status"`
	Errors []string    `json:"errors"`
	Data   interface{} `json:"data"`
}

type ApiPutRequest struct {
	Name    string      `json:"name"`
	TTL     int         `json:"ttl"`
	Payload interface{} `json:"payload"`
}
