package handler

import (
	"crud/internal/db"
	"fmt"
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
func DropUser(e echo.Context) error {
	id := e.Param("id")
	err := db.DropUserFromDB(id)
	if err != nil {
		fmt.Printf("ОШИБКА:%V", err)
	}

	log.Printf("ERROR: %v", err)

	return e.String(http.StatusOK, "success")

}
