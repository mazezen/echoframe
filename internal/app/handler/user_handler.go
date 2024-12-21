package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/echoframe/internal/app/service"
	"github.com/mazezen/echoframe/internal/models"
)

type UserHandler struct{}

// CreateUser creates a new User
func (this *UserHandler) CreateUser(c echo.Context) {
	var input models.User
	if err := c.Bind(&input); err != nil {
		panic(err)
	}

	if err := service.NewUserService().CreateUser(&input); err != nil {
		panic(err)
	}
}

// GetUser retrieves an existing User
func (this *UserHandler) GetUser(c echo.Context) {
	id := c.Param("id")

	result, err := service.NewUserService().GetUser(id)
	if err != nil {
		panic(err)
	}

	c.Set("response", result)
}

// UpdateUser updates an existing User
func (this *UserHandler) UpdateUser(c echo.Context) {
	var input models.User
	if err := c.Bind(&input); err != nil {
		panic(err)
	}

	if err := service.NewUserService().UpdateUser(&input); err != nil {
		panic(err)
	}
}

// DeleteUser deletes an existing User
func (this *UserHandler) DeleteUser(c echo.Context) {
	id := c.Param("id")

	if err := service.NewUserService().DeleteUser(id); err != nil {
		panic(err)
	}
}
