package main

import (
	"github.com/labstack/echo"
	"github.org/kevincoder/cachesystem/core"
)

func main() {
	e := echo.New()
	e.POST("/api/cache/store", core.PutCache)
	e.GET("/api/cache", core.GetCache)

	go core.RunCrons()

	e.Logger.Fatal(e.Start(":1323"))
}
