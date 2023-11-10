package router

import (
	"github.com/labstack/echo/v4"
	"github.com/ocakhasan/mongoapi/internal/controllers"
)

func Initialize(controller *controllers.PostsController) *echo.Echo {
	e := echo.New()

	api := e.Group("/api")

	api.GET("/books", controller.GetBooksWithComments())
	api.POST("/book", controller.CreateBook())
	api.GET("/author/:id/books", controller.GetAuthorBooksWithComments())

	return e
}
