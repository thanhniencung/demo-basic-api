package handler

import (
	"basic-api/repository"
	"fmt"
	"github.com/labstack/echo"
	"basic-api/model"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

/**
{
	"nanme": "Ryan1",
	"email": "abc@gmail.com"
}
 */
func (u *UserHandler) Add(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(500, model.Response{
			StatusCode: 500,
			Message:    err.Error(),
		})
	}

	err := u.UserRepo.Insert(user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(500, model.Response{
			StatusCode: 500,
			Message:     err.Error(),
		})
	}

	return c.JSON(200, model.Response{
		StatusCode: 200,
		Message:    "OK",
	})
}

