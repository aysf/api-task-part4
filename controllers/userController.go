package controllers

import (
	"aysf/day6r1/lib/database"
	"aysf/day6r1/middlewares"
	"aysf/day6r1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.String(http.StatusBadRequest, "error request")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "status ok",
		"user":    users,
	})
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetUser(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "query success",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	userInput := new(models.User)
	if err := c.Bind(userInput); err != nil {
		return err
	}
	user, err := database.CreateUser(userInput)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})

}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	userInput := new(models.User)
	if err := c.Bind(userInput); err != nil {
		return err
	}
	user, err := database.UpdateUser(id, userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully updated",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "failed to delete ",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "successfully delete data user " + c.Param("id"),
	})
}

func LoginUsersController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, err := database.LoginUsers(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to login",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    users,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	loggedUserId := middlewares.ExtractTokenUserId(c)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if loggedUserId != id {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized access !",
		})
	}

	users, err := database.GetDetailUsers((id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
