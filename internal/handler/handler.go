package handler

import (
	"crud/internal/db"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetUser(e echo.Context) error {
	id := e.Param("id")
	userData := db.GetUserFromDB(id)

	log.Printf("USERS: %v", userData)

	return e.JSON(http.StatusOK, userData)
}
