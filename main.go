package main

import (
	"crud/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.POST("/create/:id", handler.GetUser)
	/// localhost:5555/create

	err := e.Start(":5555")
	if err != nil {
		panic(err)
	}
}
