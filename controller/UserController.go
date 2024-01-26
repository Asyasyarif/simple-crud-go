package controller

import (
	"diitfin/config"
	"diitfin/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddUser(c echo.Context) error {
	u := new(model.Users)
	db := config.DB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	user := &model.Users{
		FullName: u.FullName,
		Email:    u.Email,
	}

	if err := db.Create(&user).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "success",
	}

	return c.JSON(http.StatusCreated, response)
}

func GetAllUser(c echo.Context) error {
	u := new([]model.Users)
	db := config.DB()

	if err := c.Bind(u); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := db.Find(&u).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "success",
		"data":    u,
	}

	return c.JSON(http.StatusOK, response)
}

func UserByID(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		data := map[string]interface{}{
			"message": "id is empty",
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	var u model.Users
	db := config.DB()
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data := map[string]interface{}{
				"message": "User not found",
			}
			return c.JSON(http.StatusNotFound, data)
		}

		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": u,
	}

	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	newUserData := new(model.Users)

	if id == "" {
		data := map[string]interface{}{
			"message": "id is empty",
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	if err := c.Bind(newUserData); err != nil {
		data := map[string]interface{}{
			"message": "Failed to bind request body",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	db := config.DB()
	currentUser := new(model.Users)

	if err := db.Where("id = ?", id).First(&currentUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data := map[string]interface{}{
				"message": "User not found",
			}
			return c.JSON(http.StatusNotFound, data)
		}

		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	currentUser.FullName = newUserData.FullName
	currentUser.Email = newUserData.Email

	if err := db.Save(&currentUser).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": currentUser,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		data := map[string]interface{}{
			"message": "id is empty",
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	db := config.DB()
	if err := db.Where("id = ?", id).Delete(&model.Users{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data := map[string]interface{}{
				"message": "User not found",
			}
			return c.JSON(http.StatusNotFound, data)
		}

		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "User successfully deleted",
	}
	return c.JSON(http.StatusOK, response)
}
