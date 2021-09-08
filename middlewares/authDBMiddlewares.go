package middlewares

import (
	"aysf/day6r1/config"
	"aysf/day6r1/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	tx := db.Where("email = ? AND password = ?", username, password).First(&user)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
