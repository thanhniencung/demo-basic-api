package router

import (
	"basic-api/db"
	"basic-api/handler"
	repo "basic-api/repository/repo_impl"
	"github.com/labstack/echo"
)

func Router(e *echo.Echo, sql *db.Sql) {
	handler := handler.UserHandler{
		UserRepo: repo.NewUserRepo(sql),
	}
	e.POST("/user-service/add", handler.Add)
}