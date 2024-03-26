package main

import (
	"crud/internal/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/get/:id", handler.GetUser)
	e.DELETE("/delete/:id", handler.DropUser)
	e.POST("/insert", handler.InsertUser)
	/// localhost:5555/create

	err := e.Start(":5555")
	if err != nil {
		panic(err)
	}
}
