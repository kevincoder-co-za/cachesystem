package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func PutCache(c echo.Context) error {
	var response StandardApiResponse
	var request ApiPutRequest

	err := c.Bind(&request)
	if err != nil {
		response.Status = StatusFailed
		response.Errors = []string{fmt.Sprintf("JSON body is invalid: %s", err.Error())}
	}

	ttl := time.Minute * 60

	var errors []string

	if request.TTL != 0 {
		ttl = time.Duration(request.TTL) * time.Minute
	}

	if len(request.Name) < 3 {
		errors = append(errors, "Bad cache name provided. Must be atleast 3 characters long.")
	}

	if len(errors) > 0 {
		response.Errors = errors
		response.Status = StatusFailed
		return c.JSON(http.StatusBadRequest, response)
	}

	cacheEntry := CacheEntry{
		Name:      request.Name,
		ExpiresOn: time.Now().Add(ttl),
		Payload:   request.Payload,
	}

	CacheStore.Store(request.Name, cacheEntry)
	response.Status = StatusOK
	response.Data = cacheEntry
	return c.JSON(http.StatusOK, response)

}

func GetCache(c echo.Context) error {
	var response StandardApiResponse
	name := c.QueryParam("name")

	var errors []string

	if len(name) < 3 {
		errors = append(errors, "Invalid cache name provided. Must be at least 3 characters long.")
	}

	if cache, ok := CacheStore.Load(name); ok {
		response.Data = cache
	} else {
		errors = append(errors, "No entry in the cache store found.")
	}

	if len(errors) > 0 {
		response.Errors = errors
		response.Status = StatusFailed
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Status = StatusOK

	return c.JSON(http.StatusOK, response)

}
