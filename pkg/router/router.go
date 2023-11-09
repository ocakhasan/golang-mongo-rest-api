package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ocakhasan/mongoapi/internal/controllers"
)

func Initialize(controller *controllers.PostsController) *echo.Echo {
	e := echo.New()

	api := e.Group("/api")

	api.GET("/posts", controller.GetPostsWithComments())
	api.GET("/author/:id/posts", controller.GetAuthorPostsWithComments())

	return e
}
